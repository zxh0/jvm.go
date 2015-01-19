package class

type ConstantPool struct {
    constants []Constant
}

type Constant interface{}

func (self *ConstantPool) GetConstant(index uint) (Constant) {
    // todo
    return self.constants[index]
}
