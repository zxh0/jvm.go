package class

import (
    . "jvmgo/any"
    "jvmgo/classfile"
    "jvmgo/classpath"
)

const (
    objectClassName         = "java/lang/Object"
    classClassName          = "java/lang/Class"
    threadClassName         = "java/lang/Thread"
    stringClassName         = "java/lang/String"
    serializableClassName   = "java/lang/Serializable"
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
    classPath   *classpath.ClassPath
    classMap    map[string]*Class
}

func NewClassLoader(cp *classpath.ClassPath) (*ClassLoader) {
    classMap := map[string]*Class{}
    return &ClassLoader{cp, classMap}
}

func (self *ClassLoader) Init() {
    self.LoadClass(objectClassName)
    self.LoadClass(classClassName)
    jlClassClass := self.classMap[classClassName]
    for _, class := range self.classMap {
        if class.jClass == nil {
            class.jClass = jlClassClass.NewObj()
            class.jClass.extra = class
        }
    }
    self.loadPrimitiveClasses()
    self.loadPrimitiveArrayClasses()
}

func (self *ClassLoader) loadPrimitiveClasses() {
    for primitiveType, _ := range primitiveTypes {
        self.loadPrimitiveClass(primitiveType)
    }
}
func (self *ClassLoader) loadPrimitiveClass(className string) {
    jlClassClass := self.classMap[classClassName]
    class := &Class{name: className}
    class.classLoader = self
    class.jClass = jlClassClass.NewObj()
    class.jClass.extra = class
    class.MarkInitialized()
    self.classMap[className] = class
}

func (self *ClassLoader) loadPrimitiveArrayClasses() {
    for _, arrayType := range primitiveTypes {
        if arrayType != "" {
            self.loadArrayClass(arrayType)
        }
    }
}
func (self *ClassLoader) loadArrayClass(className string) {
    jlClassClass := self.classMap[classClassName]
    jlObjecClass := self.classMap[objectClassName]

    class := &Class{name: className}
    class.classLoader = self
    class.superClass = jlObjecClass
    class.interfaces = []*Class{}
    class.jClass = jlClassClass.NewObj()
    class.jClass.extra = class
    class.MarkInitialized()
    self.classMap[className] = class
}

func (self *ClassLoader) getRefArrayClass(componentClass *Class) (*Class) {
    arrClassName := "[L" + componentClass.Name() + ";"
    return self._getRefArrayClass(arrClassName)
}
func (self *ClassLoader) _getRefArrayClass(arrClassName string) (*Class) {
    arrClass := self.classMap[arrClassName]
    if arrClass == nil {
        self.loadArrayClass(arrClassName)
    }
    return self.classMap[arrClassName]
}

// todo
func (self *ClassLoader) GetPrimitiveClass(name string) (*Class) {
    return self.getClass(name)
}

func (self *ClassLoader) StringClass() (*Class) {
    return self.getClass(stringClassName)
}
func (self *ClassLoader) ThreadClass() (*Class) {
    return self.getClass(threadClassName)
}

// todo dangerous
func (self *ClassLoader) getClass(name string) (*Class) {
    class := self.classMap[name]
    if class == nil {
        panic("class not loaded! " + name)
    }
    return class
}

func (self *ClassLoader) LoadClass(name string) (*Class) {
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

func (self *ClassLoader) reallyLoadClass(name string) (*Class) {
    //log.Printf("load: %v", name)
    class := self.parseClassFile(name)
    hackClass(class)
    self.resolveSuperClass(class)
    self.resolveInterfaces(class)
    calcStaticFieldSlots(class)
    calcInstanceFieldSlots(class)
    prepare(class)
    // todo
    class.classLoader = self
    self.classMap[name] = class

    jlClassClass := self.classMap[classClassName]
    if jlClassClass != nil {
        class.jClass = jlClassClass.NewObj()
        class.jClass.extra = class
    }

    return class
}

func (self *ClassLoader) parseClassFile(name string) (class *Class) {
    classData, err := self.classPath.ReadClassData(name)
    if err != nil {
        // todo
        panic("class not found: " + name + "!")
    }

    cf, err := classfile.ParseClassFile(classData)
    if err != nil {
        // todo
        panic("failed to parse class file: " + name + "!" + err.Error())
    }

    return newClass(cf)
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
