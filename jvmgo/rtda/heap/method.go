package heap

import (
	"fmt"
	"strings"

	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
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

func newMethod(class *Class, methodInfo *cf.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.accessFlags = methodInfo.AccessFlags()
	method.name = methodInfo.Name()
	method.descriptor = methodInfo.Descriptor()
	method.md = parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount()
	method.copyAttributes(methodInfo)
	return method
}
func (self *Method) calcArgSlotCount() {
	self.argSlotCount = self.md.argSlotCount()
	if !self.IsStatic() {
		self.argSlotCount++
	}
}
func (self *Method) copyAttributes(methodInfo *cf.MemberInfo) {
	if codeAttr := methodInfo.CodeAttribute(); codeAttr != nil {
		self.exceptions = methodInfo.ExceptionsAttribute()
		self.signature = methodInfo.Signature()
		self.code = codeAttr.Code()
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		if len(codeAttr.ExceptionTable()) > 0 {
			rtCp := self.class.constantPool
			self.copyExceptionTable(codeAttr.ExceptionTable(), rtCp)
		}
	}
	self.annotationData = methodInfo.RuntimeVisibleAnnotationsAttributeData()
	self.parameterAnnotationData = methodInfo.RuntimeVisibleParameterAnnotationsAttributeData()
	self.annotationDefaultData = methodInfo.AnnotationDefaultAttributeData()
}

func (self *Method) String() string {
	return fmt.Sprintf("{Method name:%v descriptor:%v}", self.name, self.descriptor)
}

// getters & setters
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}
func (self *Method) Slot() uint {
	return self.slot
}
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) ParameterAnnotationData() []byte {
	return self.parameterAnnotationData
}
func (self *Method) AnnotationDefaultData() []byte {
	return self.annotationDefaultData
}
func (self *Method) ParsedDescriptor() *MethodDescriptor {
	return self.md
}

func (self *Method) HackSetCode(code []byte) {
	self.code = code
}

func (self *Method) NativeMethod() interface{} {
	if self.nativeMethod == nil {
		self.nativeMethod = findNativeMethod(self)
	}
	return self.nativeMethod
}

func (self *Method) IsVoidReturnType() bool {
	return strings.HasSuffix(self.descriptor, ")V")
}

func (self *Method) isConstructor() bool {
	return !self.IsStatic() && self.name == constructorName
}
func (self *Method) IsClinit() bool {
	return self.IsStatic() &&
		self.name == clinitMethodName &&
		self.descriptor == clinitMethodDesc
}
func (self *Method) IsRegisterNatives() bool {
	return self.IsStatic() &&
		self.name == "registerNatives" &&
		self.descriptor == "()V"
}
func (self *Method) IsInitIDs() bool {
	return self.IsStatic() &&
		self.name == "initIDs" &&
		self.descriptor == "()V"
}

func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}
	if self.lineNumberTable != nil {
		return self.lineNumberTable.GetLineNumber(pc)
	}
	return -1
}
