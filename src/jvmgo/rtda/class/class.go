package class

import "jvmgo/classfile"

type Class struct {
    staticFields    []*Field
    staticMethods   []*Method
    fields          []*Field
    methods         []*Method
    // todo
}

type Field struct {
    name    string
}

type Method struct {
    name    string
}

func NewClass(cf *classfile.ClassFile) (*Class) {
    // todo
    return &Class{}
}
