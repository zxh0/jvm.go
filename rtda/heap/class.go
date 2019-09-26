package heap

import (
	"sync"

	cp "github.com/zxh0/jvm.go/classpath"
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
	staticFieldSlots   []Slot
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

func (class *Class) String() string {
	return "{Class name:" + class.name + "}"
}

// getters
func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) Name() string {
	return class.name
}
func (class *Class) Methods() []*Method {
	return class.methods
}
func (class *Class) Fields() []*Field {
	return class.fields
}
func (class *Class) StaticFieldSlots() []Slot {
	return class.staticFieldSlots
}
func (class *Class) JClass() *Object {
	return class.jClass
}
func (class *Class) SuperClass() *Class {
	return class.superClass
}
func (class *Class) Interfaces() []*Class {
	return class.interfaces
}
func (class *Class) LoadedFrom() cp.Entry {
	return class.loadedFrom
}
func (class *Class) InitCond() *sync.Cond {
	return class.initCond
}

// todo
func (class *Class) NameJlsFormat() string {
	return SlashToDot(class.name)
}

func (class *Class) InitializationNotStarted() bool {
	return class.initState < _beingInitialized // todo
}
func (class *Class) IsBeingInitialized() (bool, uintptr) {
	return class.initState == _beingInitialized, class.initThread
}
func (class *Class) IsFullyInitialized() bool {
	return class.initState == _fullyInitialized
}
func (class *Class) IsInitializationFailed() bool {
	return class.initState == _initFailed
}
func (class *Class) MarkBeingInitialized(thread uintptr) {
	class.initState = _beingInitialized
	class.initThread = thread
}
func (class *Class) MarkFullyInitialized() {
	class.initState = _fullyInitialized
}

func (class *Class) getField(name, descriptor string, isStatic bool) *Field {
	for k := class; k != nil; k = k.superClass {
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
func (class *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for k := class; k != nil; k = k.superClass {
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
func (class *Class) _getMethod(name, descriptor string, isStatic bool) *Method {
	for _, method := range class.methods {
		if method.IsStatic() == isStatic &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (class *Class) GetStaticField(name, descriptor string) *Field {
	return class.getField(name, descriptor, true)
}
func (class *Class) GetInstanceField(name, descriptor string) *Field {
	return class.getField(name, descriptor, false)
}

func (class *Class) GetStaticMethod(name, descriptor string) *Method {
	return class.getMethod(name, descriptor, true)
}
func (class *Class) GetInstanceMethod(name, descriptor string) *Method {
	return class.getMethod(name, descriptor, false)
}

func (class *Class) GetMainMethod() *Method {
	return class.GetStaticMethod(mainMethodName, mainMethodDesc)
}
func (class *Class) GetClinitMethod() *Method {
	return class._getMethod(clinitMethodName, clinitMethodDesc, true)
}

func (class *Class) NewObjWithExtra(extra interface{}) *Object {
	obj := class.NewObj()
	obj.extra = extra
	return obj
}
func (class *Class) NewObj() *Object {
	if class.instanceFieldCount > 0 {
		fields := make([]Slot, class.instanceFieldCount)
		obj := newObj(class, fields, nil)
		obj.initFields()
		return obj
	} else {
		return newObj(class, nil, nil)
	}
}
func (class *Class) NewArray(count uint) *Object {
	return NewRefArray(class, count)
}

func (class *Class) isJlObject() bool {
	return class == _jlObjectClass
}
func (class *Class) isJlCloneable() bool {
	return class == _jlCloneableClass
}
func (class *Class) isJioSerializable() bool {
	return class == _ioSerializableClass
}

// reflection
func (class *Class) GetStaticValue(fieldName, fieldDescriptor string) Slot {
	field := class.GetStaticField(fieldName, fieldDescriptor)
	return field.GetStaticValue()
}
func (class *Class) SetStaticValue(fieldName, fieldDescriptor string, value Slot) {
	field := class.GetStaticField(fieldName, fieldDescriptor)
	field.PutStaticValue(value)
}

func (class *Class) AsObj() *Object {
	return &Object{fields: class.staticFieldSlots}
}
