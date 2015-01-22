package class

import (
    "jvmgo/classfile"
    //"jvmgo/rtda"
)

const (
    cinit = "<cinit>"
    oinit = "<init>"
)

type Class struct {
    obj             Obj // todo
    constantPool    *ConstantPool
    superClassName  string
    interfaceNames  []string
    fields          []*Field
    methods         []*Method
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

func (self *Class) GetCinitMethod() (*Method) {
    return self.GetMethod(cinit)
}
func (self *Class) GetMethod(name string) (*Method) {
    for _, method := range self.methods {
        if method.name == name {
            return method
        }
    }
    // todo
    panic("method not found:" + name)
}

func (self *Class) NewObj() (*Obj) {
    // todo
    return nil
}

func newClass(cf *classfile.ClassFile) (*Class) {
    class := &Class{}
    class.constantPool = newConstantPool(cf.ConstantPool())
    class.superClassName = cf.SuperClassName()
    class.interfaceNames = cf.InterfaceNames()
    class.copyFields(cf)
    //class.fields
    //class.methods

    return class
}

func (self *Class) copyFields(cf *classfile.ClassFile) {
    cp := cf.ConstantPool()
    self.fields = make([]*Field, len(cf.Fields()))
    for i, fieldInfo := range cf.Fields() {
        self.fields[i] = newField(fieldInfo, cp, self)
    }
}

func copyMethods(cf *classfile.ClassFile) ([]Method) {
    // todo
    return nil
}
