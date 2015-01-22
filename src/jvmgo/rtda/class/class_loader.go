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

func (self *ClassLoader) getClass(name string) (*Class) {
    return self.classMap[name]
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

    class := newClass(cf)
    class.classLoader = self
    self.classMap[name] = class
    self.loadSuperClassAndInterfaces(class)

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

func NewClassLoader(cp *classpath.ClassPath) (*ClassLoader) {
    classMap := map[string]*Class{}
    return &ClassLoader{cp, classMap}
}
