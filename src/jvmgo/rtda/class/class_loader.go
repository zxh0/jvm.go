package class

import (
    "log"
    "jvmgo/classfile"
    "jvmgo/classpath"
)

type ClassLoader struct {
    classPath   *classpath.ClassPath
    classMap    map[string]*Class
}

func (self *ClassLoader) LoadClass(name string) (*Class) {
    class := self.classMap[name]
    if class != nil {
        // already loaded
        return class
    }

    class = self.reallyLoadClass(name)
    self.classMap[name] = class
    self.loadSuperClassAndInterfaces(class)

    return class
}

func (self *ClassLoader) reallyLoadClass(name string) (*Class) {
    log.Printf("load class: %v", name)
    classData, err := self.classPath.ReadClassData(name)
    if err != nil {
        // todo
        panic("class not found:" + name + "!" + err.Error())
    }

    cf, err := classfile.ParseClassFile(classData)
    if err != nil {
        // todo
        panic("failed to parse class file:" + name + "!" + err.Error())
    }

    return newClass(cf)
}

func (self *ClassLoader) loadSuperClassAndInterfaces(class *Class) {
    // todo load super class
    if class.superClassName != "" {
        self.LoadClass(class.superClassName)
    }
    // todo load interfaces
    for _, interfaceName := range class.interfaceNames {
        self.LoadClass(interfaceName)
    }
}

func NewClassLoader(cp *classpath.ClassPath) (*ClassLoader) {
    classMap := map[string]*Class{}
    return &ClassLoader{cp, classMap}
}
