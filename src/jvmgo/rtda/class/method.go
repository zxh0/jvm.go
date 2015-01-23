package class

import cf "jvmgo/classfile"

type Method struct {
    AccessFlags
    name        string
    descriptor  string
    maxStack    uint
    maxLocals   uint
    argCount    uint
    code        []byte
    class       *Class
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
func (self *Method) Class() (*Class) {
    return self.class
}
func (self *Method) Code() ([]byte) {
    return self.code
}

func newMethod(methodInfo *cf.MethodInfo, cp *cf.ConstantPool, class *Class) (*Method) {
    method := &Method{}
    method.accessFlags = methodInfo.AccessFlags()
    method.name = methodInfo.GetName(cp)
    method.descriptor = methodInfo.GetDescriptor(cp)
    if codeAttr := methodInfo.CodeAttribute(); codeAttr != nil {
        method.code = codeAttr.Code()
        method.maxStack = codeAttr.MaxStack()
        method.maxLocals = codeAttr.MaxLocals()
    }
    
    method.class = class
    return method
}

// todo
func NewStartupMethod(code []byte) (*Method) {
    method := &Method{}
    method.name = "<jvmgo>"
    method.maxStack = 8
    method.maxLocals = 8
    method.code = code
    return method
}
