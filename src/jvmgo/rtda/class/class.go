package class

import (
    "jvmgo/classfile"
    //"jvmgo/rtda"
)

type Class struct {
    Obj // todo
    superClassName  string
    interfaceNames  []string
    staticFields    []*Field
    staticMethods   []*Method
    instanceFields  []*Field
    instanceMethods []*Method
    constantPool    *ConstantPool
    classLoader     ClassLoader
    initialized     bool
    // todo
}

func (self *Class) SuperClassName() (string) {
    return self.superClassName
}
func (self *Class) IsInitialized() (bool) {
    return self.initialized
}
func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}

func GetUpmostUninitializedClassOrInterface(from *Class) (*Class) {
    if from.initialized {
        return nil
    }
    loader := from.classLoader
    if from.superClassName != "" {
        superClass := loader.getClass(from.superClassName)
        if !superClass.initialized {
            return GetUpmostUninitializedClassOrInterface(superClass)
        }
    }
    for _, interfaceName := range from.interfaceNames {
        iClass := loader.getClass(interfaceName)
        if !iClass.initialized {
            return GetUpmostUninitializedClassOrInterface(iClass)
        }
    }
    return from
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
