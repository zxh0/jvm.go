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

type Field struct {
    name string
}

func NewClass(cf *classfile.ClassFile) (*Class) {
    // todo
    return &Class{}
}
