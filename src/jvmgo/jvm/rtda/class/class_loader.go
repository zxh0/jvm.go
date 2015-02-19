package class

import (
	"fmt"
	. "jvmgo/any"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"jvmgo/jvm/options"
)

const (
	jlObjectClassName       = "java/lang/Object"
	jlClassClassName        = "java/lang/Class"
	jlCloneableClassName    = "java/lang/Cloneable"
	ioSerializableClassName = "java/io/Serializable"
	jlThreadClassName       = "java/lang/Thread"
	jlStringClassName       = "java/lang/String"
)

var (
	_jlObjectClass       *Class
	_jlClassClass        *Class
	_jlCloneableClass    *Class
	_ioSerializableClass *Class
	_jlThreadClass       *Class
	_jlStringClass       *Class
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

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	classMap := map[string]*Class{}
	return &ClassLoader{cp, classMap}
}

func (self *ClassLoader) Init() {
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
	for primitiveType, _ := range primitiveTypes {
		self.loadPrimitiveClass(primitiveType)
	}
}
func (self *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{name: className}
	class.classLoader = self
	class.jClass = _jlClassClass.NewObj()
	class.jClass.extra = class
	class.MarkInitialized()
	self.classMap[className] = class
}

func (self *ClassLoader) loadPrimitiveArrayClasses() {
	for _, descriptor := range primitiveTypes {
		self.loadArrayClass("[" + descriptor)
	}
}
func (self *ClassLoader) loadArrayClass(className string) {
	class := &Class{name: className}
	class.classLoader = self
	class.superClass = _jlObjectClass
	class.interfaces = []*Class{_jlClassClass, _ioSerializableClass}
	class.jClass = _jlClassClass.NewObj()
	class.jClass.extra = class
	createVtable(class)
	class.MarkInitialized()
	self.classMap[className] = class
}

func (self *ClassLoader) getRefArrayClass(componentClass *Class) *Class {
	arrClassName := "[L" + componentClass.Name() + ";"
	return self._getRefArrayClass(arrClassName)
}
func (self *ClassLoader) _getRefArrayClass(arrClassName string) *Class {
	arrClass := self.classMap[arrClassName]
	if arrClass == nil {
		self.loadArrayClass(arrClassName)
	}
	return self.classMap[arrClassName]
}

// todo
func (self *ClassLoader) GetPrimitiveClass(name string) *Class {
	return self.getClass(name)
}

func (self *ClassLoader) JLClassClass() *Class {
	return _jlClassClass
}
func (self *ClassLoader) StringClass() *Class {
	return _jlStringClass
}
func (self *ClassLoader) ThreadClass() *Class {
	return _jlThreadClass
}

// todo dangerous
func (self *ClassLoader) getClass(name string) *Class {
	class := self.classMap[name]
	if class == nil {
		panic("class not loaded! " + name)
	}
	return class
}

func (self *ClassLoader) LoadClass(name string) *Class {
	class := self.classMap[name]
	if class != nil {
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
	cpEntry, class := self.parseClassFile(name)
	hackClass(class)
	self.resolveSuperClass(class)
	self.resolveInterfaces(class)
	calcStaticFieldSlots(class)
	calcInstanceFieldSlots(class)
	createVtable(class)
	prepare(class)
	// todo
	class.classLoader = self
	self.classMap[name] = class

	if _jlClassClass != nil {
		class.jClass = _jlClassClass.NewObj()
		class.jClass.extra = class
	}

	if options.VerboseClass {
		fmt.Printf("[Loaded %s from %s]\n", name, cpEntry)
	}
	return class
}

func (self *ClassLoader) parseClassFile(name string) (classpath.ClassPathEntry, *Class) {
	cpEntry, classData, err := self.classPath.ReadClassData(name)
	if err != nil {
		// todo
		panic("class not found: " + name + "!")
	}

	cf, err := classfile.ParseClassFile(classData)
	if err != nil {
		// todo
		panic("failed to parse class file: " + name + "!" + err.Error())
	}

	return cpEntry, newClass(cf)
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

func calcStaticFieldSlots(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slot = slotId
			slotId++
		}
	}
	class.staticFieldCount = slotId
}

func calcInstanceFieldSlots(class *Class) {
	slotId := uint(0)
	if class.superClassName != "" {
		slotId = class.superClass.instanceFieldCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slot = slotId
			slotId++
		}
	}
	class.instanceFieldCount = slotId
}

func prepare(class *Class) {
	class.staticFieldValues = make([]Any, class.staticFieldCount)
	for _, field := range class.fields {
		if field.IsStatic() {
			class.staticFieldValues[field.slot] = field.defaultValue()
		}
	}
}
