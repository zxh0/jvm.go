package class

type ConstantFieldref struct {
    classIndex          uint16
    nameAndTypeIndex    uint16
    fieldRef            *Field
}

func (self *ConstantFieldref) Field() (*Field) {
    if self.fieldRef == nil {
        self.resolve()
    }
    return self.fieldRef
}

func (self *ConstantFieldref) resolve() {
    // todo
}
