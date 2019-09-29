package heap

import (
	"fmt"
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

type Method struct {
	ClassMember
	ExceptionTable
	MaxStack                uint
	MaxLocals               uint
	ArgSlotCount            uint
	Slot                    uint
	ParsedDescriptor        *MethodDescriptor
	Code                    []byte
	ParameterAnnotationData []byte // RuntimeVisibleParameterAnnotations_attribute
	AnnotationDefaultData   []byte // AnnotationDefault_attribute
	lineNumberTable         []classfile.LineNumberTableEntry
	exIndexTable            []uint16    // TODO: rename
	nativeMethod            interface{} // cannot use package 'native' because of cycle import!
	Instructions            interface{} // []instructions.Instruction
}

func newMethod(class *Class, cf *classfile.ClassFile, methodInfo classfile.MemberInfo) *Method {
	method := &Method{}
	method.Class = class
	method.AccessFlags = AccessFlags(methodInfo.AccessFlags)
	method.Name = cf.GetUTF8(methodInfo.NameIndex)
	method.Descriptor = cf.GetUTF8(methodInfo.DescriptorIndex)
	method.ParsedDescriptor = parseMethodDescriptor(method.Descriptor)
	method.calcArgSlotCount()
	method.copyAttributes(cf, methodInfo)
	return method
}
func (method *Method) calcArgSlotCount() {
	method.ArgSlotCount = method.ParsedDescriptor.argSlotCount()
	if !method.IsStatic() {
		method.ArgSlotCount++
	}
}
func (method *Method) copyAttributes(cf *classfile.ClassFile, methodInfo classfile.MemberInfo) {
	if codeAttr := methodInfo.GetCodeAttribute(); codeAttr != nil {
		method.exIndexTable = methodInfo.GetExceptionIndexTable()
		method.Signature = cf.GetUTF8(methodInfo.GetSignatureIndex())
		method.Code = codeAttr.Code
		method.MaxStack = uint(codeAttr.MaxStack)
		method.MaxLocals = uint(codeAttr.MaxLocals)
		method.lineNumberTable = codeAttr.GetLineNumberTable()
		if len(codeAttr.ExceptionTable) > 0 {
			rtCp := method.Class.ConstantPool
			method.copyExceptionTable(codeAttr.ExceptionTable, rtCp)
		}
	}
	method.AnnotationData = methodInfo.GetRuntimeVisibleAnnotationsAttributeData()
	method.ParameterAnnotationData = methodInfo.GetRuntimeVisibleParameterAnnotationsAttributeData()
	method.AnnotationDefaultData = methodInfo.GetAnnotationDefaultAttributeData()
}

func (method *Method) String() string {
	return fmt.Sprintf("{Method name:%v descriptor:%v}", method.Name, method.Descriptor)
}

func (method *Method) NativeMethod() interface{} {
	if method.nativeMethod == nil {
		method.nativeMethod = findNativeMethod(method)
	}
	return method.nativeMethod
}

func (method *Method) IsVoidReturnType() bool {
	return strings.HasSuffix(method.Descriptor, ")V")
}

func (method *Method) isConstructor() bool {
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
