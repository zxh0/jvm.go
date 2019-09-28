package heap

import (
	"fmt"

	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/vmerrors"
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
	verbose   bool
}

func BootLoader() *ClassLoader {
	return bootLoader
}

func InitBootLoader(cp *classpath.ClassPath, verbose bool) {
	bootLoader = &ClassLoader{
		classPath: cp,
		classMap:  map[string]*Class{},
		verbose:   verbose,
	}
	bootLoader._init()
}

func (loader *ClassLoader) _init() {
	_jlObjectClass = loader.LoadClass(jlObjectClassName)
	_jlClassClass = loader.LoadClass(jlClassClassName)
	for _, class := range loader.classMap {
		if class.jClass == nil {
			class.jClass = _jlClassClass.NewObj()
			class.jClass.extra = class
		}
	}
	_jlCloneableClass = loader.LoadClass(jlCloneableClassName)
	_ioSerializableClass = loader.LoadClass(ioSerializableClassName)
	_jlThreadClass = loader.LoadClass(jlThreadClassName)
	_jlStringClass = loader.LoadClass(jlStringClassName)
	loader.loadPrimitiveClasses()
	loader.loadPrimitiveArrayClasses()
}

func (loader *ClassLoader) loadPrimitiveClasses() {
	for _, primitiveType := range PrimitiveTypes {
		loader.loadPrimitiveClass(primitiveType.Name)
	}
}
func (loader *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{name: className}
	//class.classLoader = loader
	class.jClass = _jlClassClass.NewObj()
	class.jClass.extra = class
	class.MarkFullyInitialized()
	loader.classMap[className] = class
}

func (loader *ClassLoader) loadPrimitiveArrayClasses() {
	for _, primitiveType := range PrimitiveTypes {
		loader.loadArrayClass(primitiveType.ArrayClassName)
	}
}
func (loader *ClassLoader) loadArrayClass(className string) *Class {
	class := &Class{name: className}
	//class.classLoader = loader
	class.superClass = _jlObjectClass
	class.interfaces = []*Class{_jlCloneableClass, _ioSerializableClass}
	class.jClass = _jlClassClass.NewObj()
	class.jClass.extra = class
	createVtable(class)
	class.MarkFullyInitialized()
	loader.classMap[className] = class
	return class
}

func (loader *ClassLoader) getRefArrayClass(componentClass *Class) *Class {
	arrClassName := "[L" + componentClass.Name() + ";"
	return loader._getRefArrayClass(arrClassName)
}
func (loader *ClassLoader) _getRefArrayClass(arrClassName string) *Class {
	if arrClass, ok := loader.classMap[arrClassName]; ok {
		return arrClass
	}
	return loader.loadArrayClass(arrClassName)
}

func (loader *ClassLoader) ClassPath() *classpath.ClassPath {
	return loader.classPath
}

func (loader *ClassLoader) JLObjectClass() *Class {
	return _jlObjectClass
}
func (loader *ClassLoader) JLClassClass() *Class {
	return _jlClassClass
}
func (loader *ClassLoader) JLStringClass() *Class {
	return _jlStringClass
}
func (loader *ClassLoader) JLThreadClass() *Class {
	return _jlThreadClass
}

// todo
func (loader *ClassLoader) GetPrimitiveClass(name string) *Class {
	return loader.getClass(name)
}

func (loader *ClassLoader) FindLoadedClass(name string) *Class {
	if class, ok := loader.classMap[name]; ok {
		return class
	}
	return nil
}

// todo dangerous
func (loader *ClassLoader) getClass(name string) *Class {
	if class, ok := loader.classMap[name]; ok {
		return class
	}
	panic("class not loaded! " + name)
}

func (loader *ClassLoader) LoadClass(name string) *Class {
	if class, ok := loader.classMap[name]; ok {
		// already loaded
		return class
	} else if name[0] == '[' {
		// array class
		return loader._getRefArrayClass(name)
	} else {
		return loader.reallyLoadClass(name)
	}
}

func (loader *ClassLoader) reallyLoadClass(name string) *Class {
	cpEntry, data := loader.readClassData(name)
	class := loader._loadClass(name, data)
	class.loadedFrom = cpEntry

	if loader.verbose {
		fmt.Printf("[Loaded %s from %s]\n", name, cpEntry)
	}

	return class
}

func (loader *ClassLoader) readClassData(name string) (classpath.Entry, []byte) {
	cpEntry, classData, err := loader.classPath.ReadClass(name)
	if err != nil {
		panic(vmerrors.NewClassNotFoundError(SlashToDot(name)))
	}

	return cpEntry, classData
}

func (loader *ClassLoader) parseClassData(name string, data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		// todo
		panic("failed to parse class file: " + name + "! " + err.Error())
	}

	return newClass(cf)
}

func (loader *ClassLoader) _loadClass(name string, data []byte) *Class {
	class := loader.parseClassData(name, data)
	hackClass(class)
	loader.resolveSuperClass(class)
	loader.resolveInterfaces(class)
	calcStaticFieldSlotIds(class)
	calcInstanceFieldSlotIds(class)
	createVtable(class)
	prepare(class)
	// todo
	//class.classLoader = loader
	loader.classMap[name] = class

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
func (loader *ClassLoader) resolveSuperClass(class *Class) {
	if class.superClassName != "" {
		class.superClass = loader.LoadClass(class.superClassName)
	}
}
func (loader *ClassLoader) resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = loader.LoadClass(interfaceName)
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
	class.staticFieldSlots = make([]Slot, class.staticFieldCount)
	for _, field := range class.fields {
		if field.IsStatic() {
			class.staticFieldSlots[field.slotId] = EmptySlot // TODO
		}
	}
}

// todo
func (loader *ClassLoader) DefineClass(name string, data []byte) *Class {
	return loader._loadClass(name, data)
}
