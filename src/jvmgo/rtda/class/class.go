package class

import (
    //"fmt"
    . "jvmgo/any"
)

const (
    _initializing = 1
    _initialized  = 2
)

type Class struct {
    constantPool        *ConstantPool
    name                string // thisClassName
    superClassName      string
    interfaceNames      []string
    fields              []*Field
    methods             []*Method
    staticFieldCount    uint
    instanceFieldCount  uint
    staticFieldValues   []Any
    state               int
    jClass              *Obj // java.lang.Class instance
    superClass          *Class
    classLoader         *ClassLoader // defining class loader
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
    return self.GetStaticMethod(mainMethodName, mainMethodDesc)
}
func (self *Class) GetClinitMethod() (*Method) {
    return self.GetStaticMethod(clinitMethodName, clinitMethodDesc)
}
func (self *Class) GetStaticMethod(name, descriptor string) (*Method) {
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
        obj := &Obj{self, fields}
        obj.zeroFields()
        return obj
    } else {
        return &Obj{self, nil}
    }
}

func (self *Class) zeroStaticFields() {
    for _, f := range self.fields {
        if f.IsStatic() {
            self.staticFieldValues[f.slot] = f.zeroValue()
        }
    }
}
