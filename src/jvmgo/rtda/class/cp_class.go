package class

type ConstantClass struct {
    name    uint16
    cp      *ConstantPool
    class   *Class
}

func (self *ConstantClass) Class() (*Class) {
    if self.class == nil {
        self.resolve()
    }
    return self.class
}

func (self *ConstantClass) resolve() {
    // todo
    // methodArea?
    //ma := self.cp.class.classMap
}
