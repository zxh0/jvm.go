package class

type ConstantMethodref struct {
    classIndex          uint16
    nameAndTypeIndex    uint16
    methodRef           *Method
}

func (self *ConstantMethodref) Method() (*Method) {
    if self.methodRef == nil {
        self.resolve()
    }
    return self.methodRef
}

func (self *ConstantMethodref) resolve() {
    // todo
}
