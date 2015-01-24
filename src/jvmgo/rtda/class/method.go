package class

import cf "jvmgo/classfile"

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
    maxStack    uint
    maxLocals   uint
    argCount    uint
    code        []byte
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

func (self *Method) IsClinit() (bool) {
    return self.name == clinitMethodName && self.descriptor == clinitMethodDesc
}
func (self *Method) IsRegisterNatives() (bool) {
    return self.name == registerNativesMethodName &&
            self.descriptor == registerNativesMethodDesc
}

func newMethod(class *Class, methodInfo *cf.MethodInfo) (*Method) {
    method := &Method{}
    method.class = class
    method.accessFlags = methodInfo.AccessFlags()
    method.name = methodInfo.GetName()
    method.descriptor = methodInfo.GetDescriptor()
    if codeAttr := methodInfo.CodeAttribute(); codeAttr != nil {
        method.code = codeAttr.Code()
        method.maxStack = codeAttr.MaxStack()
        method.maxLocals = codeAttr.MaxLocals()
    }
    
    return method
}

// todo
func NewStartupMethod(code []byte) (*Method) {
    method := &Method{}
    method.class = &Class{name:"~jvmgo"}
    method.name = "<jvmgo>"
    method.maxStack = 8
    method.maxLocals = 8
    method.code = code
    return method
}
