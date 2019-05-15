package heap

import (
	"sync"

	cp "github.com/zxh0/jvm.go/jvmgo/classpath"
)

// initialization state
const (
	_notInitialized   = 0 // This Class object is verified and prepared but not initialized.
	_beingInitialized = 1 // This Class object is being initialized by some particular thread T.
	_fullyInitialized = 2 // This Class object is fully initialized and ready for use.
	_initFailed       = 3 // This Class object is in an erroneous state, perhaps because initialization was attempted and failed.
)

// name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
type Class struct {
	AccessFlags
	ClassAttributes
	constantPool       *ConstantPool
	name               string // thisClassName
	superClassName     string
	interfaceNames     []string
	fields             []*Field
	methods            []*Method
	instanceFieldCount uint
	staticFieldCount   uint
	staticFieldSlots   []interface{}
	vtable             []*Method // virtual method table
	jClass             *Object   // java.lang.Class instance
	superClass         *Class
	interfaces         []*Class
	loadedFrom         cp.Entry // todo
	initState          int
	initCond           *sync.Cond
	initThread         uintptr
	//classLoader        *ClassLoader      // defining class loader
}

func (self *Class) String() string {
	return "{Class name:" + self.name + "}"
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Name() string {
	return self.name
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) StaticFieldSlots() []interface{} {
	return self.staticFieldSlots
}
func (self *Class) JClass() *Object {
	return self.jClass
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) Interfaces() []*Class {
	return self.interfaces
}
func (self *Class) LoadedFrom() cp.Entry {
	return self.loadedFrom
}
func (self *Class) InitCond() *sync.Cond {
	return self.initCond
}

// todo
func (self *Class) NameJlsFormat() string {
	return SlashToDot(self.name)
}

func (self *Class) InitializationNotStarted() bool {
	return self.initState < _beingInitialized // todo
}
func (self *Class) IsBeingInitialized() (bool, uintptr) {
	return self.initState == _beingInitialized, self.initThread
}
func (self *Class) IsFullyInitialized() bool {
	return self.initState == _fullyInitialized
}
func (self *Class) IsInitializationFailed() bool {
	return self.initState == _initFailed
}
func (self *Class) MarkBeingInitialized(thread uintptr) {
	self.initState = _beingInitialized
	self.initThread = thread
}
func (self *Class) MarkFullyInitialized() {
	self.initState = _fullyInitialized
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
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
func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
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

// todo
func (self *Class) _getMethod(name, descriptor string, isStatic bool) *Method {
	for _, method := range self.methods {
		if method.IsStatic() == isStatic &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) GetStaticField(name, descriptor string) *Field {
	return self.getField(name, descriptor, true)
}
func (self *Class) GetInstanceField(name, descriptor string) *Field {
	return self.getField(name, descriptor, false)
}

func (self *Class) GetStaticMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, true)
}
func (self *Class) GetInstanceMethod(name, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

func (self *Class) GetMainMethod() *Method {
	return self.GetStaticMethod(mainMethodName, mainMethodDesc)
}
func (self *Class) GetClinitMethod() *Method {
	return self._getMethod(clinitMethodName, clinitMethodDesc, true)
}

func (self *Class) NewObjWithExtra(extra interface{}) *Object {
	obj := self.NewObj()
	obj.extra = extra
	return obj
}
func (self *Class) NewObj() *Object {
	if self.instanceFieldCount > 0 {
		fields := make([]interface{}, self.instanceFieldCount)
		obj := newObj(self, fields, nil)
		obj.initFields()
		return obj
	} else {
		return newObj(self, nil, nil)
	}
}
func (self *Class) NewArray(count uint) *Object {
	return NewRefArray(self, count)
}

func (self *Class) isJlObject() bool {
	return self == _jlObjectClass
}
func (self *Class) isJlCloneable() bool {
	return self == _jlCloneableClass
}
func (self *Class) isJioSerializable() bool {
	return self == _ioSerializableClass
}

// reflection
func (self *Class) GetStaticValue(fieldName, fieldDescriptor string) interface{} {
	field := self.GetStaticField(fieldName, fieldDescriptor)
	return field.GetStaticValue()
}
func (self *Class) SetStaticValue(fieldName, fieldDescriptor string, value interface{}) {
	field := self.GetStaticField(fieldName, fieldDescriptor)
	field.PutStaticValue(value)
}

func (self *Class) AsObj() *Object {
	return &Object{fields: self.staticFieldSlots}
}
