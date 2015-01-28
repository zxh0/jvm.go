package class

import (
    . "jvmgo/any"
    cf "jvmgo/classfile"
)

const (
    mainMethodName              = "main"
    mainMethodDesc              = "([Ljava/lang/String;)V"
    clinitMethodName            = "<clinit>"
    clinitMethodDesc            = "()V"
    registerNativesMethodName   = "registerNatives"
    registerNativesMethodDesc   = "()V"
)

type Method struct {
    ClassMember
    ExceptionTable
    maxStack        uint
    maxLocals       uint
    argCount        uint
    code            []byte
    nativeMethod    Any // cannot use package 'native' because of cycle import!
}

// getters
func (self *Method) MaxStack() (uint) {
    return self.maxStack
}
func (self *Method) MaxLocals() (uint) {
    return self.maxLocals
}
func (self *Method) ArgCount() (uint) {
    return self.argCount
}
func (self *Method) Code() ([]byte) {
    return self.code
}
func (self *Method) NativeMethod() (Any) {
    return self.nativeMethod
}

func (self *Method) IsClinit() (bool) {
    return self.name == clinitMethodName && self.descriptor == clinitMethodDesc
}
func (self *Method) IsRegisterNatives() (bool) {
    return self.IsStatic() &&
            self.name == registerNativesMethodName &&
            self.descriptor == registerNativesMethodDesc

}

func newMethod(class *Class, methodInfo *cf.MethodInfo) (*Method) {
    method := &Method{}
    method.class = class
    method.accessFlags = methodInfo.AccessFlags()
    method.name = methodInfo.GetName()
    method.descriptor = methodInfo.GetDescriptor()
    method.argCount = methodInfo.GetArgCount()
    // if !method.IsStatic() { // todo
    //     method.argCount++
    // }
    if codeAttr := methodInfo.CodeAttribute(); codeAttr != nil {
        method.code = codeAttr.Code()
        method.maxStack = codeAttr.MaxStack()
        method.maxLocals = codeAttr.MaxLocals()
        if len(codeAttr.ExceptionTable()) > 0 {
            method.copyExceptionTable(codeAttr.ExceptionTable())
        }
    }
    
    return method
}

// hack
func NewStartupMethod(code []byte, classLoader Any) (*Method) {
    method := &Method{}
    method.class = &Class{name:"~jvmgo", classLoader:classLoader.(*ClassLoader)}
    method.name = "<jvmgo>"
    method.maxStack = 8
    method.maxLocals = 8
    method.code = code
    return method
}
