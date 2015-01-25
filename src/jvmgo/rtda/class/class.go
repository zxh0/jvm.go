package class

import (
    //"fmt"
    . "jvmgo/any"
    //"jvmgo/classfile"
    //"jvmgo/rtda"
)

type Class struct {
    obj                 *Obj // todo
    constantPool        *ConstantPool
    name                string
    superClassName      string
    interfaceNames      []string
    fields              []*Field
    methods             []*Method
    classLoader         *ClassLoader
    instanceFieldCount  uint
    initialized         bool
    // todo
}

func (self *Class) Obj() (*Obj) {
    return self.obj
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
func (self *Class) NotInitialized() (bool) {
    return !self.initialized
}
func (self *Class) MarkInitialized() {
    self.initialized = true
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
