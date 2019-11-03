package heap

import (
	"strings"

	"github.com/zxh0/jvm.go/classfile"
)

const (
	mainMethodName   = "main"
	mainMethodDesc   = "([Ljava/lang/String;)V"
	clinitMethodName = "<clinit>"
	clinitMethodDesc = "()V"
	constructorName  = "<init>"
)

type MethodData struct {
	MaxStack                uint
	MaxLocals               uint
	Code                    []byte
	exceptionTable          []classfile.ExceptionTableEntry
	lineNumberTable         []classfile.LineNumberTableEntry
	ParameterAnnotationData []byte // RuntimeVisibleParameterAnnotations_attribute
	AnnotationDefaultData   []byte // AnnotationDefault_attribute
}

type Method struct {
	ClassMember
	MethodData
	MethodDescriptor
	ParamCount     uint
	ParamSlotCount uint
	Slot           uint
	exIndexTable   []uint16    // TODO: rename
	Instructions   interface{} // []instructions.Instruction
}

func newMethod(class *Class, cf *classfile.ClassFile, cfMember classfile.MemberInfo, slot uint) *Method {
	method := &Method{}
	method.Class = class
	method.Slot = slot
	method.copyMemberData(cf, cfMember)
	method.copyAttributes(cfMember)
	method.parseDescriptor()
	return method
}

func (method *Method) copyAttributes(cfMember classfile.MemberInfo) {
	if codeAttr, found := cfMember.GetCodeAttribute(); found {
		method.exIndexTable = cfMember.GetExceptionIndexTable()
		method.MaxStack = uint(codeAttr.MaxStack)
		method.MaxLocals = uint(codeAttr.MaxLocals)
		method.Code = codeAttr.Code
		method.exceptionTable = codeAttr.ExceptionTable
		method.lineNumberTable = codeAttr.GetLineNumberTable()
	}
	method.ParameterAnnotationData = cfMember.GetRuntimeVisibleParameterAnnotationsAttributeData()
	method.AnnotationDefaultData = cfMember.GetAnnotationDefaultAttributeData()
}

func (method *Method) parseDescriptor() {
	method.MethodDescriptor = parseMethodDescriptor(method.Descriptor)
	method.ParamCount = uint(len(method.ParameterTypes))
	method.ParamSlotCount = method.getParamSlotCount()
	if !method.IsStatic() {
		method.ParamSlotCount++
	}
}

func (method *Method) IsVoidReturnType() bool {
	return strings.HasSuffix(method.Descriptor, ")V")
}

func (method *Method) IsConstructor() bool {
	return !method.IsStatic() && method.Name == constructorName
}
func (method *Method) IsClinit() bool {
	return method.IsStatic() &&
		method.Name == clinitMethodName &&
		method.Descriptor == clinitMethodDesc
}
func (method *Method) IsRegisterNatives() bool {
	return method.IsStatic() &&
		method.Name == "registerNatives" &&
		method.Descriptor == "()V"
}
func (method *Method) IsInitIDs() bool {
	return method.IsStatic() &&
		method.Name == "initIDs" &&
		method.Descriptor == "()V"
}

func (method *Method) FindExceptionHandler(exClass *Class, pc int) int {
	for _, handler := range method.exceptionTable {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= int(handler.StartPc) && pc < int(handler.EndPc) {
			if handler.CatchType == 0 {
				// catch all
				return int(handler.HandlerPc)
			}

			catchType := method.Class.GetConstantClass(uint(handler.CatchType))
			if catchType.GetClass() == exClass ||
				catchType.GetClass().isSuperClassOf(exClass) {

				return int(handler.HandlerPc)
			}
		}
	}
	return -1
}

func (method *Method) GetLineNumber(pc int) int {
	if method.IsNative() {
		return -2
	}
	for i := len(method.lineNumberTable) - 1; i >= 0; i-- {
		entry := method.lineNumberTable[i]
		if pc >= int(entry.StartPC) {
			return int(entry.LineNumber)
		}
	}
	return -1
}
