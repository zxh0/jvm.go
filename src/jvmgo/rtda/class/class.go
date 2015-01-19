package class

import "jvmgo/classfile"

type Class struct {
    staticFields    []*Field
    staticMethods   []*Method
    fields          []*Field
    methods         []*Method
    constantPool    *ConstantPool
    // todo
}

func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}

type Field struct {
    name string
}

func NewClass(cf *classfile.ClassFile) (*Class) {
    // todo
    return &Class{}
}
