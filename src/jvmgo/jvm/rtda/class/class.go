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

// name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
type Class struct {
    AccessFlags
    constantPool        *ConstantPool
    sourceFile          string
    name                string // thisClassName
    superClassName      string
    interfaceNames      []string
    fields              []*Field
    methods             []*Method
    staticFieldCount    uint
    instanceFieldCount  uint
    staticFieldValues   []Any
    jClass              *Obj // java.lang.Class instance
    superClass          *Class
    interfaces          []*Class
    classLoader         *ClassLoader // defining class loader
    state               int
}

func (self *Class) String() string {
    return "{Class name:" + self.name + "}"
}

// getters
func (self *Class) ConstantPool() (*ConstantPool) {
    return self.constantPool
}
func (self *Class) SourceFile() string {
    return self.sourceFile
}
func (self *Class) Name() string {
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

func (self *Class) InitializationNotStarted() (bool) {
    return self.state < _initializing // todo
}
func (self *Class) MarkInitializing() {
    self.state = _initializing
}
func (self *Class) MarkInitialized() {
    self.state = _initialized
}

func (self *Class) getField(name, descriptor string, isStatic bool) (*Field) {
    for k := self; k != nil; k = k.superClass {
        for _, field := range k.fields {
            if field.IsStatic() == isStatic &&
                    field.name == name &&
                    field.descriptor == descriptor {

                return field
            }
        }
    }
    // todo
    return nil
}
func (self *Class) getMethod(name, descriptor string, isStatic bool) (*Method) {
    for k := self; k != nil; k = k.superClass {
        for _, method := range k.methods {
            if method.IsStatic() == isStatic &&
                    method.name == name &&
                    method.descriptor == descriptor {

                return method
            }
        }
    }
    // todo
    return nil
}

func (self *Class) GetStaticField(name, descriptor string) (*Field) {
    return self.getField(name, descriptor, true)
}
func (self *Class) GetInstanceField(name, descriptor string) (*Field) {
    return self.getField(name, descriptor, false)
}

func (self *Class) GetMainMethod() (*Method) {
    return self.GetStaticMethod(mainMethodName, mainMethodDesc)
}
func (self *Class) GetClinitMethod() (*Method) {
    return self.GetStaticMethod(clinitMethodName, clinitMethodDesc)
}
func (self *Class) GetStaticMethod(name, descriptor string) (*Method) {
    return self.getMethod(name, descriptor, true)
}
func (self *Class) GetInstanceMethod(name, descriptor string) (*Method) {
    return self.getMethod(name, descriptor, false)
}

func (self *Class) getArrayClass() (*Class) {
    return self.classLoader.getRefArrayClass(self)
}

func (self *Class) NewObjWithExtra(extra Any) (*Obj) {
    obj := self.NewObj()
    obj.extra = extra
    return obj
}
func (self *Class) NewObj() (*Obj) {
    if self.instanceFieldCount > 0 {
        fields := make([]Any, self.instanceFieldCount)
        obj := &Obj{self, fields, nil}
        obj.initFields()
        return obj
    } else {
        return &Obj{self, nil, nil}
    }
}
func (self *Class) NewArray(count uint) (*Obj) {
    return NewRefArray(self, count)
}

func (self *Class) isObject() bool {
    return self.name == jlObjectName
}

func (self *Class) isSubClassOf(c *Class) bool {
    for k := self.superClass; k != nil; k = k.superClass {
        if k == c {
            return true
        }
    }
    return false
}
func (self *Class) isSuperClassOf(c *Class) bool {
    return c.isSubClassOf(self)
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

func (self *Class) AsObj() (*Obj) {
    return &Obj{fields: self.staticFieldValues}
}
