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
    jlStringName = "java/lang/String"
)

type ClassLoader struct {
    classPath   *classpath.ClassPath
    classMap    map[string]*Class
}

func (self *ClassLoader) Init() {
    self.LoadClass(jlObjectName)
    self.LoadClass(jlClassName)
    jlClassClass := self.classMap[jlClassName]
    for _, class := range self.classMap {
        if class.jClass == nil {
            class.jClass = jlClassClass.NewObj()
        }
    }
}

// todo dangerous
func (self *ClassLoader) getClass(name string) (*Class) {
    class := self.classMap[name]
    if class == nil {
        panic("class not loaded! " + name)
    }
    return class
}

func (self *ClassLoader) StringClass() (*Class) {
    return self.LoadClass(jlStringName)
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
    class.classLoader = self
    self.loadSuperClassAndInterfaces(class)
    calcStaticFieldSlots(class)
    calcInstanceFieldSlots(class)
    self.classMap[name] = class
    class.staticFieldValues = make([]Any, class.staticFieldCount)
    class.zeroStaticFields()

    jlClassClass := self.classMap[jlClassName]
    if jlClassClass != nil {
        class.jClass = jlClassClass.NewObj()
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
func (self *ClassLoader) loadSuperClassAndInterfaces(class *Class) {
    if class.superClassName != "" {
        class.superClass = self.LoadClass(class.superClassName)
    }
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

func NewClassLoader(cp *classpath.ClassPath) (*ClassLoader) {
    classMap := map[string]*Class{}
    return &ClassLoader{cp, classMap}
}
