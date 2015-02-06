package class

import (
    //"fmt"
    . "jvmgo/any"
    //cf "jvmgo/classfile"
)

const (
    _initializing = 1
    _initialized  = 2
)

type Class struct {
    constantPool        *ConstantPool
    AccessFlags
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
    interfaces          []*Class
    classLoader         *ClassLoader // defining class loader
}

func (self *Class) String() string {
    return "{Class name:" + self.name + "}"
}

// getters
func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}
func (self *Class) Name() (string) {
    return self.name
}
func (self *Class) JClass() (*Obj) {
    return self.jClass
}
func (self *Class) SuperClass() (*Class) {
    return self.superClass
}
func (self *Class) Interfaces() ([]*Class) {
    return self.interfaces
}
func (self *Class) ClassLoader() (*ClassLoader) {
    return self.classLoader
}

func (self *Class) IsPrimitive() (bool) {
    return isPrimitiveType(self.name)
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

func (self *Class) GetStaticField(name, descriptor string) (*Field) {
    return self.getField(name, descriptor, true)
}
func (self *Class) GetInstanceField(name, descriptor string) (*Field) {
    return self.getField(name, descriptor, false)
}
func (self *Class) getField(name, descriptor string, isStatic bool) (*Field) {
    for _, field := range self.fields {
        if field.IsStatic() == isStatic &&
                field.name == name &&
                field.descriptor == descriptor {

            return field
        }
    }
    // todo
    return nil
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

func (self *Class) getArrayClass() (*Class) {
    return self.classLoader.getRefArrayClass(self)
}

func (self *Class) NewObj() (*Obj) {
    if self.instanceFieldCount > 0 {
        fields := make([]Any, self.instanceFieldCount)
        obj := &Obj{self, fields, nil}
        obj.zeroFields()
        return obj
    } else {
        return &Obj{self, nil, nil}
    }
}
func (self *Class) NewArray(count int) (*Obj) {
    return NewRefArray(self, count)
}

// reflection
func (self *Class) GetStaticValue(fieldName, fieldDescriptor string) Any {
    field := self.GetStaticField(fieldName, fieldDescriptor)
    return field.GetStaticValue()
}
func (self *Class) SetStaticValue(fieldName, fieldDescriptor string, value Any) {
    field := self.GetStaticField(fieldName, fieldDescriptor)
    field.PutStaticValue(value)
}
