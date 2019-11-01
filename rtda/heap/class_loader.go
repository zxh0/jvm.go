package heap

import (
	"fmt"

	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/vm"
	"github.com/zxh0/jvm.go/vmutils"
)

const (
	jlObjectClassName       = "java/lang/Object"
	jlClassClassName        = "java/lang/Class"
	jlStringClassName       = "java/lang/String"
	jlThreadClassName       = "java/lang/Thread"
	jlCloneableClassName    = "java/lang/Cloneable"
	ioSerializableClassName = "java/io/Serializable"
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
	rt        *Runtime
	classPath *classpath.ClassPath
	classMap  map[string]*Class // loaded classes
	verbose   bool
	// some frequently used classes
	jlObjectClass       *Class
	jlClassClass        *Class
	jlStringClass       *Class
	jlThreadClass       *Class
	jlCloneableClass    *Class
	ioSerializableClass *Class
}

func newBootLoader(cp *classpath.ClassPath, verbose bool) *ClassLoader {
	return &ClassLoader{
		classPath: cp,
		classMap:  map[string]*Class{},
		verbose:   verbose,
	}
}

func (loader *ClassLoader) init() {
	loader.jlObjectClass = loader.LoadClass(jlObjectClassName)
	loader.jlClassClass = loader.LoadClass(jlClassClassName)
	for _, class := range loader.classMap {
		if class.JClass == nil {
			class.JClass = loader.jlClassClass.NewObj()
			class.JClass.Extra = class
		}
	}
	loader.jlCloneableClass = loader.LoadClass(jlCloneableClassName)
	loader.ioSerializableClass = loader.LoadClass(ioSerializableClassName)
	loader.jlThreadClass = loader.LoadClass(jlThreadClassName)
	loader.jlStringClass = loader.LoadClass(jlStringClassName)
	loader.loadPrimitiveClasses()
	loader.loadPrimitiveArrayClasses()
}

func (loader *ClassLoader) loadPrimitiveClasses() {
	for _, primitiveType := range primitiveTypes {
		loader.loadPrimitiveClass(primitiveType.Name)
	}
}
func (loader *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{Name: className}
	class.bootLoader = loader
	class.JClass = loader.jlClassClass.NewObj()
	class.JClass.Extra = class
	class.MarkFullyInitialized()
	loader.classMap[className] = class
}

func (loader *ClassLoader) loadPrimitiveArrayClasses() {
	for _, primitiveType := range primitiveTypes {
		loader.loadArrayClass(primitiveType.ArrayClassName)
	}
}
func (loader *ClassLoader) loadArrayClass(className string) *Class {
	class := &Class{Name: className}
	class.bootLoader = loader
	class.SuperClass = loader.jlObjectClass
	class.Interfaces = []*Class{loader.jlCloneableClass, loader.ioSerializableClass}
	class.JClass = loader.jlClassClass.NewObj()
	class.JClass.Extra = class
	createVtable(class)
	class.MarkFullyInitialized()
	loader.classMap[className] = class
	return class
}

func (loader *ClassLoader) getRefArrayClass(componentClass *Class) *Class {
	arrClassName := "[L" + componentClass.Name + ";"
	return loader.getRefArrayClassByName(arrClassName)
}
func (loader *ClassLoader) getRefArrayClassByName(arrClassName string) *Class {
	if arrClass, ok := loader.classMap[arrClassName]; ok {
		return arrClass
	}
	return loader.loadArrayClass(arrClassName)
}

func (loader *ClassLoader) JLObjectClass() *Class {
	return loader.jlObjectClass
}
func (loader *ClassLoader) JLClassClass() *Class {
	return loader.jlClassClass
}
func (loader *ClassLoader) JLStringClass() *Class {
	return loader.jlStringClass
}
func (loader *ClassLoader) JLThreadClass() *Class {
	return loader.jlThreadClass
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
		return loader.getRefArrayClassByName(name)
	} else {
		return loader.reallyLoadClass(name)
	}
}

func (loader *ClassLoader) reallyLoadClass(name string) *Class {
	cpEntry, data := loader.readClassData(name)
	class := loader.loadClass(name, data)
	class.LoadedFrom = cpEntry

	if loader.verbose {
		fmt.Printf("[Loaded %s from %s]\n", name, cpEntry)
	}

	return class
}

func (loader *ClassLoader) readClassData(name string) (classpath.Entry, []byte) {
	cpEntry, classData := loader.classPath.ReadClass(name)
	if classData == nil {
		panic(vm.NewClassNotFoundError(vmutils.SlashToDot(name)))
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

func (loader *ClassLoader) loadClass(name string, data []byte) *Class {
	class := loader.parseClassData(name, data)
	hackClass(class)
	loader.resolveSuperClass(class)
	loader.resolveInterfaces(class)
	calcStaticFieldSlotIds(class)
	calcInstanceFieldSlotIds(class)
	createVtable(class)
	prepare(class)
	// todo
	class.bootLoader = loader
	loader.classMap[name] = class

	if loader.jlClassClass != nil {
		class.JClass = loader.jlClassClass.NewObj()
		class.JClass.Extra = class
	}

	return class
}

// todo
func hackClass(class *Class) {
	if class.Name == "java/lang/ClassLoader" {
		loadLibrary := class.GetStaticMethod("loadLibrary", "(Ljava/lang/Class;Ljava/lang/String;Z)V")
		loadLibrary.Code = []byte{0xb1} // return void
	}
}

// todo
func (loader *ClassLoader) resolveSuperClass(class *Class) {
	if class.superClassName != "" {
		class.SuperClass = loader.LoadClass(class.superClassName)
	}
}
func (loader *ClassLoader) resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.Interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.Interfaces[i] = loader.LoadClass(interfaceName)
		}
	}
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.Fields {
		if field.IsStatic() {
			field.SlotId = slotId
			slotId++
		}
	}
	class.staticFieldCount = slotId
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClassName != "" {
		slotId = class.SuperClass.instanceFieldCount
	}
	for _, field := range class.Fields {
		if !field.IsStatic() {
			field.SlotId = slotId
			slotId++
		}
	}
	class.instanceFieldCount = slotId
}

func prepare(class *Class) {
	class.StaticFieldSlots = make([]Slot, class.staticFieldCount)
	for _, field := range class.Fields {
		if field.IsStatic() {
			class.StaticFieldSlots[field.SlotId] = EmptySlot // TODO
		}
	}
}

// todo
func (loader *ClassLoader) DefineClass(name string, data []byte) *Class {
	return loader.loadClass(name, data)
}
