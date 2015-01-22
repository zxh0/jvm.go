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
    classLoader     ClassLoader
    initialized     bool
    // todo
}

func (self *Class) IsInitialized() (bool) {
    return self.initialized
}
func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}

func (self *Class) NewObj() (*Obj) {
    // todo
    return nil
}

func newClass(cf *classfile.ClassFile) (*Class) {
    cfCp := cf.ConstantPool()
    rtCp := newConstantPool(cfCp)

    // todo
    // copy consts
    class := &Class{}
    class.constantPool = rtCp
    return class
}
