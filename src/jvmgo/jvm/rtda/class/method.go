package class

import (
	"fmt"
	. "jvmgo/any"
	cf "jvmgo/classfile"
	"jvmgo/util"
	"strings"
)

const (
	mainMethodName            = "main"
	mainMethodDesc            = "([Ljava/lang/String;)V"
	clinitMethodName          = "<clinit>"
	clinitMethodDesc          = "()V"
	constructorName           = "<init>"
	registerNativesMethodName = "registerNatives"
	registerNativesMethodDesc = "()V"
)

type Method struct {
	ClassMember
	ExceptionTable
	maxStack        uint
	maxLocals       uint
	argCount        uint
	md              *MethodDescriptor
	code            []byte
	lineNumberTable *cf.LineNumberTableAttribute
	exceptions      *cf.ExceptionsAttribute
	nativeMethod    Any // cannot use package 'native' because of cycle import!
}

func newMethod(class *Class, methodInfo *cf.MethodInfo) *Method {
	method := &Method{}
	method.class = class
	method.accessFlags = methodInfo.AccessFlags()
	method.name = methodInfo.Name()
	method.descriptor = methodInfo.Descriptor()
	method.md = parseMethodDescriptor(method.descriptor)
	method.argCount = method.md.argCount()
	method.copyAttributes(methodInfo)
	return method
}
func (self *Method) copyAttributes(methodInfo *cf.MethodInfo) {
	if codeAttr := methodInfo.CodeAttribute(); codeAttr != nil {
		self.code = codeAttr.Code()
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		self.exceptions = methodInfo.ExceptionsAttribute()
		self.signature = methodInfo.Signature()
		if len(codeAttr.ExceptionTable()) > 0 {
			rtCp := self.class.constantPool
			self.copyExceptionTable(codeAttr.ExceptionTable(), rtCp)
		}
	}
	if rvaAttr := methodInfo.RuntimeVisibleAnnotationsAttribute(); rvaAttr != nil {
		self.annotationData = util.CastUint8sToInt8s(rvaAttr.Info())
	}
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
func (self *Method) ArgCount() uint {
	return self.argCount
}
func (self *Method) ParsedDescriptor() *MethodDescriptor {
	return self.md
}
func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) HackSetCode(code []byte) {
	self.code = code
}

func (self *Method) NativeMethod() Any {
	if self.nativeMethod == nil {
		self.nativeMethod = findNativeMethod(self)
	}
	return self.nativeMethod
}

// argCount for static method
// argCount+1 for instance method
func (self *Method) ActualArgCount() uint {
	if self.IsStatic() {
		return self.argCount
	} else {
		return self.argCount + 1
	}
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
		self.name == registerNativesMethodName &&
		self.descriptor == registerNativesMethodDesc
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
