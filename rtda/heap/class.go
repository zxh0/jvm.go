package heap

import (
	"sync"

	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/module"
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
	SuperClass         *Class
	Interfaces         []*Class
	instanceFieldCount uint
	staticFieldCount   uint
	StaticFieldSlots   []Slot
	vtable             []*Method // virtual method table
	JClass             *Object   // java.lang.Class instance
	LoadedFrom         module.Module
	bootLoader         *ClassLoader // TODO
	initState          int
	InitCond           *sync.Cond
	initThread         uintptr
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

func (class *Class) AsObj() *Object {
	return &Object{Fields: class.StaticFieldSlots}
}
