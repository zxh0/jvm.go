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

func (self *ConstantClass) resolve() {
    classMap := self.cp.class.classMap
    class := classMap.getClass(self.name)
    if class != nil {
        self.class = class
    } else {
        self.loadClass()   
    }
}

// todo
func (self *ConstantClass) loadClass() {
    loader := self.cp.class.classLoader
    class := loader.LoadClass(self.name)
    self.cp.class.classMap.putClass(self.name, class)
    self.class = class
}
