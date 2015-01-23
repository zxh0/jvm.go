package class

import (
    //"fmt"
    "jvmgo/classfile"
    //"jvmgo/rtda"
)

const (
    class_init = "<clinit>"
    obj_init = "<init>"
)

type Class struct {
    obj             Obj // todo
    constantPool    *ConstantPool
    name            string
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
func (self *Class) Name() (string) {
    return self.name
}
func (self *Class) IsInitialized() (bool) {
    return self.initialized
}
func (self *Class) MarkInitialized() {
    self.initialized = true
}

func (self *Class) GetClinitMethod() (*Method) {
    return self.GetMethod(class_init)
}
func (self *Class) GetMethod(name string) (*Method) {
    for _, method := range self.methods {
        if method.name == name {
            return method
        }
    }
    // todo
    return nil
}

func (self *Class) NewObj() (*Obj) {
    // todo
    return nil
}

func newClass(cf *classfile.ClassFile) (*Class) {
    class := &Class{}
    class.constantPool = newConstantPool(cf.ConstantPool())
    class.name = cf.ClassName()
    class.superClassName = cf.SuperClassName()
    class.interfaceNames = cf.InterfaceNames()
    class.copyFields(cf)
    class.copyMethods(cf)
    return class
}

func (self *Class) copyFields(cf *classfile.ClassFile) {
    cp := cf.ConstantPool()
    self.fields = make([]*Field, len(cf.Fields()))
    for i, fieldInfo := range cf.Fields() {
        self.fields[i] = newField(fieldInfo, cp, self)
    }
}

func (self *Class) copyMethods(cf *classfile.ClassFile) {
    cp := cf.ConstantPool()
    self.methods = make([]*Method, len(cf.Methods()))
    for i, methodInfo := range cf.Methods() {
        self.methods[i] = newMethod(methodInfo, cp, self)
    }
}
