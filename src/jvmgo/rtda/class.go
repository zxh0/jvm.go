package rtda

import "jvmgo/classfile"
import "jvmgo/instructions"

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
    bcr     *instructions.BytecodeReader
}

func NewClass(cf *classfile.ClassFile) (*Class) {
    // todo
    return &Class{}
}
