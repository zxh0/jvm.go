package class

type ConstantClass struct {
    name    string
    cp      *ConstantPool
    class   *Class
}

func (self *ConstantClass) Class() (*Class) {
    if self.class == nil {
        self.resolve()
    }
    return self.class
}

// todo
func (self *ConstantClass) resolve() {
    // load class
    loader := self.cp.class.classLoader
    self.class = loader.loadClass(self.name)
}
