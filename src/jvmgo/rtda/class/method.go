package class

type Method struct {
    //name    string
    maxStack    uint16
    maxLocals   uint16
    class       *Class
    code        []byte
}

// getters
func (self *Method) MaxStack() (uint16) {
    return self.maxStack
}
func (self *Method) MaxLocals() (uint16) {
    return self.maxLocals
}
func (self *Method) Class() (*Class) {
    return self.class
}
func (self *Method) Code() ([]byte) {
    return self.code
}
