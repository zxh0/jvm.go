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

func (self ClassLoader) LoadClass(name string) (*Class) {
    class := self.classMap[name]
    if class != nil {
        return class
    }

    log.Printf("load class: %v", name)
    classData, err := self.classPath.ReadClassData(name)
    if err != nil {
        // todo
        panic("class not found:" + name)
    }

    cf, err := classfile.ParseClassFile(classData)
    if err != nil {
        // todo
        panic("failed to parse class file:" + name)
    }

    class = newClass(cf)
    self.classMap[name] = class

    // todo load super class
    if class.superClassName != "" {
        self.LoadClass(class.superClassName)
    }

    return class
}

func NewClassLoader(cp *classpath.ClassPath) (*ClassLoader) {
    classMap := map[string]*Class{}
    return &ClassLoader{cp, classMap}
}
