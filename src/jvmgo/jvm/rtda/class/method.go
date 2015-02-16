package class

import (
	"fmt"
	. "jvmgo/any"
	cf "jvmgo/classfile"
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
	code            []byte
	lineNumberTable *cf.LineNumberTableAttribute
	nativeMethod    Any // cannot use package 'native' because of cycle import!
	md              *MethodDescriptor
}

func newMethod(class *Class, methodInfo *cf.MethodInfo) *Method {
	method := &Method{}
	method.class = class
	method.accessFlags = methodInfo.AccessFlags()
	method.name = methodInfo.Name()
	method.descriptor = methodInfo.Descriptor()
	if codeAttr := methodInfo.CodeAttribute(); codeAttr != nil {
		method.code = codeAttr.Code()
		method.maxStack = codeAttr.MaxStack()
		method.maxLocals = codeAttr.MaxLocals()
		method.lineNumberTable = codeAttr.LineNumberTableAttribute()
		rtCp := method.class.constantPool
		if len(codeAttr.ExceptionTable()) > 0 {
			method.copyExceptionTable(codeAttr.ExceptionTable(), rtCp)
		}
	}

	method.md = parseMethodDescriptor(method.descriptor)
	method.argCount = method.md.argCount()
	return method
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
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) SetCode(code []byte) {
	self.code = code
}
func (self *Method) NativeMethod() Any {
	return self.nativeMethod
}
func (self *Method) MethodDescriptor() *MethodDescriptor {
	return self.md
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
