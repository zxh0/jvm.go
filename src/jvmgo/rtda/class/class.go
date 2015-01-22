package class

import (
    "jvmgo/classfile"
    //"jvmgo/rtda"
)

type Class struct {
    obj             Obj // todo
    superClassName  string
    interfaceNames  []string
    fields          []*Field
    methods         []*Method
    constantPool    *ConstantPool
    classLoader     *ClassLoader
    initialized     bool
    // todo
}

// func (self *Class) SuperClassName() (string) {
//     return self.superClassName
// }
func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}
func (self *Class) IsInitialized() (bool) {
    return self.initialized
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

    class.superClassName = cf.SuperClassName()
    class.interfaceNames = cf.InterfaceNames()


    return class
}
