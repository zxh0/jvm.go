package class

import (
    //"fmt"
    . "jvmgo/any"
    cf "jvmgo/classfile"
)

const (
    _initializing = 1
    _initialized  = 2
)

type Class struct {
    constantPool        *ConstantPool
    cf.AccessFlags
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

func (self *Class) GetFields(publicOnly bool) ([]*Field) {
    if publicOnly {
        publicFields := make([]*Field, 0, len(self.fields))
        for _, field := range self.fields {
            if field.IsPublic() {
                n := len(publicFields)
                publicFields := publicFields[:n + 1]
                publicFields[n] = field
            }
        }
        return publicFields
    } else {
        return self.fields
    }
}
func (self *Class) GetMethods(publicOnly bool) ([]*Method) {
    result := make([]*Method, 0, len(self.methods))
    for _, method := range self.methods {
        if !method.IsClinit() && !method.isConstructor() {
            if !publicOnly || method.IsPublic() {
                n := len(result)
                result := result[:n + 1]
                result[n] = method
            }
        }
    }
    return result
}
func (self *Class) GetConstructors(publicOnly bool) ([]*Method) {
    constructors := make([]*Method, 0, len(self.methods))
    for _, method := range self.methods {
        if method.isConstructor() {
            if !publicOnly || method.IsPublic() {
                n := len(constructors)
                constructors := constructors[:n + 1]
                constructors[n] = method
            }
        }
    }
    return constructors
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

func (self *Class) GetConstructor(descriptor string) (*Method) {
    c := self.GetMethod(constructorName, descriptor)
    if c != nil && !c.IsStatic() {
        return c
    } else {
        return nil
    }
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
