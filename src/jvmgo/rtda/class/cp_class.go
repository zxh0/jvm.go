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
    // todo
    classMap := self.cp.class.classMap
    class := classMap.getClass(self.name)
    if class != nil {
        self.class = class
    } else {
        // load link init
    }
}
