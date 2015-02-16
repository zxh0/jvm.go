package class

var _shimClass = &Class{name:"~shim"}
var _returnCode = []byte{0xb1} // return

func NewReturnMethod(maxStack, maxLocals uint) (*Method) {
    method := &Method{}
    method.class = _shimClass
    method.name = "<return>"
    method.accessFlags = ACC_STATIC
    method.maxStack = maxStack
    method.maxLocals = maxLocals
    method.code = _returnCode
    return method
}

func NewAthrowMethod() (*Method) {
    // todo
    return nil
}
