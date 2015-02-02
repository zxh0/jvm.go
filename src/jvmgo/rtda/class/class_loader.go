package class

import (
    //"fmt"
    //"log"
    . "jvmgo/any"
    "jvmgo/classfile"
    "jvmgo/classpath"
)

const (
    jlObjectName = "java/lang/Object"
    jlClassName = "java/lang/Class"
    jlThreadName = "java/lang/Thread"
    jlStringName = "java/lang/String"
)

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
    self.LoadClass(jlObjectName)
    self.LoadClass(jlClassName)
    jlClassClass := self.classMap[jlClassName]
    for _, class := range self.classMap {
        if class.jClass == nil {
            class.jClass = jlClassClass.NewObj()
            class.jClass.extra = class
        }
    }
    self.loadPrimitiveClasses()
}

func (self *ClassLoader) loadPrimitiveClasses() {
    primitiveTypes := []string{"boolean", "byte", "char", "short", "int", "long", "float", "double"}
    for _, primitiveType := range primitiveTypes {
        self.loadPrimitiveClass(primitiveType)
    }
}

func (self *ClassLoader) loadPrimitiveClass(name string) {
    jlClassClass := self.classMap[jlClassName]
    class := &Class{name: name}
    class.jClass = jlClassClass.NewObj()
    class.jClass.extra = class
    self.classMap[name] = class
}

func (self *ClassLoader) StringClass() (*Class) {
    return self.getClass(jlStringName)
}
func (self *ClassLoader) ThreadClass() (*Class) {
    return self.getClass(jlThreadName)
}

// todo
func (self *ClassLoader) GetPrimitiveClass(name string) (*Class) {
    return self.getClass(name)
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
    } else {
        return self.reallyLoadClass(name)
    }
}

func (self *ClassLoader) reallyLoadClass(name string) (*Class) {
    //log.Printf("load: %v", name)
    class := self.parseClassFile(name)
    self.resolveSuperClass(class)
    self.resolveInterfaces(class)
    calcStaticFieldSlots(class)
    calcInstanceFieldSlots(class)
    prepare(class)
    // todo
    class.classLoader = self
    self.classMap[name] = class

    jlClassClass := self.classMap[jlClassName]
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

    return cf2class(cf)
}

// todo
func (self *ClassLoader) resolveSuperClass(class *Class) {
    if class.superClassName != "" {
        class.superClass = self.LoadClass(class.superClassName)
    }
}
func (self *ClassLoader) resolveInterfaces(class *Class) {
    for _, interfaceName := range class.interfaceNames {
        self.LoadClass(interfaceName)
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
