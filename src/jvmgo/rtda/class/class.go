package class

import (
    //"fmt"
    . "jvmgo/any"
    //"jvmgo/classfile"
    //"jvmgo/rtda"
)

const (
    _initializing = 1
    _initialized  = 2
)

type Class struct {
    obj                 *Obj // static fields live here
    jClass              *Obj // java.lang.Class instance
    constantPool        *ConstantPool
    name                string
    superClassName      string
    interfaceNames      []string
    fields              []*Field
    methods             []*Method
    classLoader         *ClassLoader
    staticFieldCount    uint
    instanceFieldCount  uint
    state               int
    // todo
}

// func (self *Class) Fields() ([]*Field){
//     return self.fields
// }

// todo
func (self *Class) Obj() (*Obj) {
    return self.obj
}
func (self *Class) JClass() (*Obj) {
    return self.jClass
}
func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}
func (self *Class) Name() (string) {
    return self.name
}
func (self *Class) ClassLoader() (*ClassLoader) {
    return self.classLoader
}

func (self *Class) InitializationNotStarted() (bool) {
    return self.state < _initializing // todo
}
func (self *Class) MarkInitializing() {
    self.state = _initializing
}
func (self *Class) MarkInitialized() {
    self.state = _initialized
}

func (self *Class) GetField(name, descriptor string) (*Field) {
    for _, field := range self.fields {
        if field.name == name && field.descriptor == descriptor {
            return field
        }
    }
    // todo
    return nil
}

func (self *Class) GetMainMethod() (*Method) {
    return self.getStaticMethod(mainMethodName, mainMethodDesc)
}
func (self *Class) GetClinitMethod() (*Method) {
    return self.getStaticMethod(clinitMethodName, clinitMethodDesc)
}
func (self *Class) getStaticMethod(name, descriptor string) (*Method) {
    method := self.GetMethod(name, descriptor)
    if method != nil && method.IsStatic() {
        return method
    } else {
        return nil
    }
}
func (self *Class) GetMethod(name, descriptor string) (*Method) {
    for _, method := range self.methods {
        if method.name == name && method.descriptor == descriptor {
            return method
        }
    }
    // todo
    return nil
}

func (self *Class) NewObj() (*Obj) {
    if self.instanceFieldCount > 0 {
        fields := make([]Any, self.instanceFieldCount)
        return &Obj{self, fields}
    } else {
        return &Obj{self, nil}
    }
}

// func newClass() (*Class) {
//     class := &Class{}
//     class.obj = &Obj{class:class} 
//     return class
// }
