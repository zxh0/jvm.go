package class

//import cf "jvmgo/classfile"

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
    self.class = loader.LoadClass(self.name)
}
