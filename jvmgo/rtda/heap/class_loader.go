package heap

import (
	"fmt"

	"github.com/zxh0/jvm.go/jvmgo/classfile"
	"github.com/zxh0/jvm.go/jvmgo/classpath"
	"github.com/zxh0/jvm.go/jvmgo/jerrors"
	"github.com/zxh0/jvm.go/jvmgo/options"
)

const (
	jlObjectClassName       = "java/lang/Object"
	jlClassClassName        = "java/lang/Class"
	jlStringClassName       = "java/lang/String"
	jlThreadClassName       = "java/lang/Thread"
	jlCloneableClassName    = "java/lang/Cloneable"
	ioSerializableClassName = "java/io/Serializable"
)

var (
	bootLoader           *ClassLoader // bootstrap class loader
	_jlObjectClass       *Class
	_jlClassClass        *Class
	_jlStringClass       *Class
	_jlThreadClass       *Class
	_jlCloneableClass    *Class
	_ioSerializableClass *Class
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/

// the bootstrap class loader
type ClassLoader struct {
	classPath *classpath.ClassPath
	classMap  map[string]*Class
}

func BootLoader() *ClassLoader {
	return bootLoader
}

func InitBootLoader(cp *classpath.ClassPath) {
	bootLoader = &ClassLoader{
		classPath: cp,
		classMap:  map[string]*Class{},
	}
	bootLoader._init()
}

func (self *ClassLoader) _init() {
	_jlObjectClass = self.LoadClass(jlObjectClassName)
	_jlClassClass = self.LoadClass(jlClassClassName)
	for _, class := range self.classMap {
		if class.jClass == nil {
			class.jClass = _jlClassClass.NewObj()
			class.jClass.extra = class
		}
	}
	_jlCloneableClass = self.LoadClass(jlCloneableClassName)
	_ioSerializableClass = self.LoadClass(ioSerializableClassName)
	_jlThreadClass = self.LoadClass(jlThreadClassName)
	_jlStringClass = self.LoadClass(jlStringClassName)
	self.loadPrimitiveClasses()
	self.loadPrimitiveArrayClasses()
}

func (self *ClassLoader) loadPrimitiveClasses() {
	for _, primitiveType := range PrimitiveTypes {
		self.loadPrimitiveClass(primitiveType.Name)
	}
}
func (self *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{name: className}
	//class.classLoader = self
	class.jClass = _jlClassClass.NewObj()
	class.jClass.extra = class
	class.MarkFullyInitialized()
	self.classMap[className] = class
}

func (self *ClassLoader) loadPrimitiveArrayClasses() {
	for _, primitiveType := range PrimitiveTypes {
		self.loadArrayClass(primitiveType.ArrayClassName)
	}
}
func (self *ClassLoader) loadArrayClass(className string) *Class {
	class := &Class{name: className}
	//class.classLoader = self
	class.superClass = _jlObjectClass
	class.interfaces = []*Class{_jlCloneableClass, _ioSerializableClass}
	class.jClass = _jlClassClass.NewObj()
	class.jClass.extra = class
	createVtable(class)
	class.MarkFullyInitialized()
	self.classMap[className] = class
	return class
}

func (self *ClassLoader) getRefArrayClass(componentClass *Class) *Class {
	arrClassName := "[L" + componentClass.Name() + ";"
	return self._getRefArrayClass(arrClassName)
}
func (self *ClassLoader) _getRefArrayClass(arrClassName string) *Class {
	if arrClass, ok := self.classMap[arrClassName]; ok {
		return arrClass
	}
	return self.loadArrayClass(arrClassName)
}

func (self *ClassLoader) ClassPath() *classpath.ClassPath {
	return self.classPath
}

func (self *ClassLoader) JLObjectClass() *Class {
	return _jlObjectClass
}
func (self *ClassLoader) JLClassClass() *Class {
	return _jlClassClass
}
func (self *ClassLoader) JLStringClass() *Class {
	return _jlStringClass
}
func (self *ClassLoader) JLThreadClass() *Class {
	return _jlThreadClass
}

// todo
func (self *ClassLoader) GetPrimitiveClass(name string) *Class {
	return self.getClass(name)
}

func (self *ClassLoader) FindLoadedClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	return nil
}

// todo dangerous
func (self *ClassLoader) getClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class
	}
	panic("class not loaded! " + name)
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		// already loaded
		return class
	} else if name[0] == '[' {
		// array class
		return self._getRefArrayClass(name)
	} else {
		return self.reallyLoadClass(name)
	}
}

func (self *ClassLoader) reallyLoadClass(name string) *Class {
	cpEntry, data := self.readClassData(name)
	class := self._loadClass(name, data)
	class.loadedFrom = cpEntry

	if options.VerboseClass {
		fmt.Printf("[Loaded %s from %s]\n", name, cpEntry)
	}

	return class
}

func (self *ClassLoader) readClassData(name string) (classpath.Entry, []byte) {
	cpEntry, classData, err := self.classPath.ReadClass(name)
	if err != nil {
		panic(jerrors.NewClassNotFoundError(SlashToDot(name)))
	}

	return cpEntry, classData
}

func (self *ClassLoader) parseClassData(name string, data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		// todo
		panic("failed to parse class file: " + name + "!" + err.Error())
	}

	return newClass(cf)
}

func (self *ClassLoader) _loadClass(name string, data []byte) *Class {
	class := self.parseClassData(name, data)
	hackClass(class)
	self.resolveSuperClass(class)
	self.resolveInterfaces(class)
	calcStaticFieldSlotIds(class)
	calcInstanceFieldSlotIds(class)
	createVtable(class)
	prepare(class)
	// todo
	//class.classLoader = self
	self.classMap[name] = class

	if _jlClassClass != nil {
		class.jClass = _jlClassClass.NewObj()
		class.jClass.extra = class
	}

	return class
}

// todo
func hackClass(class *Class) {
	if class.name == "java/lang/ClassLoader" {
		loadLibrary := class.GetStaticMethod("loadLibrary", "(Ljava/lang/Class;Ljava/lang/String;Z)V")
		loadLibrary.code = []byte{0xb1} // return void
	}
}

// todo
func (self *ClassLoader) resolveSuperClass(class *Class) {
	if class.superClassName != "" {
		class.superClass = self.LoadClass(class.superClassName)
	}
}
func (self *ClassLoader) resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = self.LoadClass(interfaceName)
		}
	}
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
		}
	}
	class.staticFieldCount = slotId
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClassName != "" {
		slotId = class.superClass.instanceFieldCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
		}
	}
	class.instanceFieldCount = slotId
}

func prepare(class *Class) {
	class.staticFieldSlots = make([]interface{}, class.staticFieldCount)
	for _, field := range class.fields {
		if field.IsStatic() {
			class.staticFieldSlots[field.slotId] = field.defaultValue()
		}
	}
}

// todo
func (self *ClassLoader) DefineClass(name string, data []byte) *Class {
	return self._loadClass(name, data)
}
