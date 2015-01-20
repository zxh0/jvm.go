package class

type Method struct {
    name    string
    class   *Class
    code    []byte
}

// getters
func (self *Method) Class() (*Class) {
    return self.class
}
func (self *Method) Code() ([]byte) {
    return self.code
}
