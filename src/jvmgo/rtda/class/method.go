package class

type Method struct {
    name    string
    class   *Class
}

func (self *Method) Class() (*Class) {
    return self.class
}
