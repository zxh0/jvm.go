package heap

import (
	"fmt"
	"strings"

	cf "github.com/zxh0/jvm.go/classfile"
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
	maxStack                uint
	maxLocals               uint
	argSlotCount            uint
	slot                    uint
	md                      *MethodDescriptor
	code                    []byte
	parameterAnnotationData []byte // RuntimeVisibleParameterAnnotations_attribute
	annotationDefaultData   []byte // AnnotationDefault_attribute
	lineNumberTable         *cf.LineNumberTableAttribute
	exceptions              *cf.ExceptionsAttribute
	nativeMethod            interface{} // cannot use package 'native' because of cycle import!
	Instructions            interface{} // []instructions.Instruction
}

func newMethod(class *Class, methodInfo cf.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.accessFlags = methodInfo.AccessFlags
	method.name = methodInfo.Name()
	method.descriptor = methodInfo.Descriptor()
	method.md = parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount()
	method.copyAttributes(methodInfo)
	return method
}
func (method *Method) calcArgSlotCount() {
	method.argSlotCount = method.md.argSlotCount()
	if !method.IsStatic() {
		method.argSlotCount++
	}
}
func (method *Method) copyAttributes(methodInfo cf.MemberInfo) {
	if codeAttr := methodInfo.CodeAttribute(); codeAttr != nil {
		method.exceptions = methodInfo.ExceptionsAttribute()
		method.signature = methodInfo.Signature()
		method.code = codeAttr.Code
		method.maxStack = uint(codeAttr.MaxStack)
		method.maxLocals = uint(codeAttr.MaxLocals)
		method.lineNumberTable = codeAttr.LineNumberTableAttribute()
		if len(codeAttr.ExceptionTable) > 0 {
			rtCp := method.class.constantPool
			method.copyExceptionTable(codeAttr.ExceptionTable, rtCp)
		}
	}
	method.annotationData = methodInfo.RuntimeVisibleAnnotationsAttributeData()
	method.parameterAnnotationData = methodInfo.RuntimeVisibleParameterAnnotationsAttributeData()
	method.annotationDefaultData = methodInfo.AnnotationDefaultAttributeData()
}

func (method *Method) String() string {
	return fmt.Sprintf("{Method name:%v descriptor:%v}", method.name, method.descriptor)
}

// getters & setters
func (method *Method) MaxStack() uint {
	return method.maxStack
}
func (method *Method) MaxLocals() uint {
	return method.maxLocals
}
func (method *Method) ArgSlotCount() uint {
	return method.argSlotCount
}
func (method *Method) Slot() uint {
	return method.slot
}
func (method *Method) Code() []byte {
	return method.code
}
func (method *Method) ParameterAnnotationData() []byte {
	return method.parameterAnnotationData
}
func (method *Method) AnnotationDefaultData() []byte {
	return method.annotationDefaultData
}
func (method *Method) ParsedDescriptor() *MethodDescriptor {
	return method.md
}

func (method *Method) HackSetCode(code []byte) {
	method.code = code
}

func (method *Method) NativeMethod() interface{} {
	if method.nativeMethod == nil {
		method.nativeMethod = findNativeMethod(method)
	}
	return method.nativeMethod
}

func (method *Method) IsVoidReturnType() bool {
	return strings.HasSuffix(method.descriptor, ")V")
}

func (method *Method) isConstructor() bool {
	return !method.IsStatic() && method.name == constructorName
}
func (method *Method) IsClinit() bool {
	return method.IsStatic() &&
		method.name == clinitMethodName &&
		method.descriptor == clinitMethodDesc
}
func (method *Method) IsRegisterNatives() bool {
	return method.IsStatic() &&
		method.name == "registerNatives" &&
		method.descriptor == "()V"
}
func (method *Method) IsInitIDs() bool {
	return method.IsStatic() &&
		method.name == "initIDs" &&
		method.descriptor == "()V"
}

func (method *Method) GetLineNumber(pc int) int {
	if method.IsNative() {
		return -2
	}
	if method.lineNumberTable != nil {
		return method.lineNumberTable.GetLineNumber(pc)
	}
	return -1
}
