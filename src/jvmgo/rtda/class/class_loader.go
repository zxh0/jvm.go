package class

import (
    //"fmt"
    //"log"
    "jvmgo/classfile"
    "jvmgo/classpath"
)

type ClassLoader struct {
    classPath   *classpath.ClassPath
    classMap    map[string]*Class
}

// todo GetClass
func (self *ClassLoader) getClass(name string) (*Class) {
    // todo
    return self.classMap[name]
}

func (self *ClassLoader) StringClass() (*Class) {
    return self.LoadClass("java/lang/String")
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

    class := cf2class(cf)
    if class.name == "java/lang/Class" {
        class.obj.class = class
    } else {
        class.obj.class = self.classMap["java/lang/Class"]
    }
    class.classLoader = self
    self.classMap[name] = class
    self.loadSuperClassAndInterfaces(class)
    self.initInstanceFields(class)

    return class
}

// todo
func (self *ClassLoader) loadSuperClassAndInterfaces(class *Class) {
    if class.superClassName != "" {
        self.LoadClass(class.superClassName)
    }
    for _, interfaceName := range class.interfaceNames {
        self.LoadClass(interfaceName)
    }
}

func (self *ClassLoader) initInstanceFields(class *Class) {
    slotId := uint(0)
    if class.superClassName != "" {
        superClass := self.getClass(class.superClassName)
        slotId = superClass.instanceFieldCount
    }
    for _, field := range class.fields {
        field.slot = slotId
        slotId++
    }
    class.instanceFieldCount = slotId
}

func NewClassLoader(cp *classpath.ClassPath) (*ClassLoader) {
    classMap := map[string]*Class{}
    return &ClassLoader{cp, classMap}
}
