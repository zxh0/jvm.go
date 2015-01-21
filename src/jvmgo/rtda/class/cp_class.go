package class

type ConstantClass struct {
    nameIndex   uint16
    class       *Class
}

func (self *ConstantClass) Class() (*Class) {
    if self.class == nil {
        self.resolve()
    }
    return self.class
}

func (self *ConstantClass) resolve() {
    // todo
}
