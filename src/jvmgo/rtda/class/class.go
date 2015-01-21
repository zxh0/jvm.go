package class

import (
    "jvmgo/classfile"
    //"jvmgo/rtda"
)

type Class struct {
    Obj // todo
    staticFields    []*Field
    staticMethods   []*Method
    instanceFields  []*Field
    instanceMethods []*Method
    constantPool    *ConstantPool
    classMap        *ClassMap
    // todo
}

func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}

func (self *Class) NewObj() (*Obj) {
    // todo
    return nil
}

func NewClass(cf *classfile.ClassFile) (*Class) {
    // todo
    // copy consts
    return &Class{}
}
