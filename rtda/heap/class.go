package heap

import (
	"sync"

	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/vmutils"
)

// initialization state
const (
	_notInitialized   = 0 // This Class object is verified and prepared but not initialized.
	_beingInitialized = 1 // This Class object is being initialized by some particular thread T.
	_fullyInitialized = 2 // This Class object is fully initialized and ready for use.
	_initFailed       = 3 // This Class object is in an erroneous state, perhaps because initialization was attempted and failed.
)

type EnclosingMethod struct {
	ClassName        string
	MethodName       string
	MethodDescriptor string
}

// name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
type Class struct {
	classfile.AccessFlags
	ConstantPool
	Name               string // thisClassName
	superClassName     string
	interfaceNames     []string
	SourceFile         string
	Signature          string
	AnnotationData     []byte // RuntimeVisibleAnnotations_attribute
	EnclosingMethod    *EnclosingMethod
	Fields             []*Field
	Methods            []*Method
	instanceFieldCount uint
	staticFieldCount   uint
	StaticFieldSlots   []Slot
	vtable             []*Method // virtual method table
	JClass             *Object   // java.lang.Class instance
	SuperClass         *Class
	Interfaces         []*Class
	LoadedFrom         classpath.Entry
	initState          int
	InitCond           *sync.Cond
	initThread         uintptr
	bootLoader         *ClassLoader // TODO
}

func (class *Class) String() string {
	return "{Class name:" + class.Name + "}"
}

// todo
func (class *Class) NameJlsFormat() string {
	return vmutils.SlashToDot(class.Name)
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
	for k := class; k != nil; k = k.SuperClass {
		for _, field := range k.Fields {
			if field.IsStatic() == isStatic &&
				field.Name == name &&
				field.Descriptor == descriptor {

				return field
			}
		}
	}
	// todo
	return nil
}
func (class *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for k := class; k != nil; k = k.SuperClass {
		for _, method := range k.Methods {
			if method.IsStatic() == isStatic &&
				method.Name == name &&
				method.Descriptor == descriptor {

				return method
			}
		}
	}
	// todo
	return nil
}

func (class *Class) getDeclaredMethod(name, descriptor string, isStatic bool) *Method {
	for _, method := range class.Methods {
		if method.IsStatic() == isStatic &&
			method.Name == name &&
			method.Descriptor == descriptor {

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
	return class.getDeclaredMethod(clinitMethodName, clinitMethodDesc, true)
}

func (class *Class) NewObjWithExtra(extra interface{}) *Object {
	obj := class.NewObj()
	obj.Extra = extra
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
	return newRefArray(class, count)
}

func (class *Class) isJlObject() bool {
	return class == class.bootLoader.jlObjectClass
}
func (class *Class) isJlCloneable() bool {
	return class == class.bootLoader.jlCloneableClass
}
func (class *Class) isJioSerializable() bool {
	return class == class.bootLoader.ioSerializableClass
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
	return &Object{Fields: class.StaticFieldSlots}
}
