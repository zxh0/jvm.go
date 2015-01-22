package class

import (
    "jvmgo/classfile"
    "jvmgo/classpath"
)

type ClassLoader struct {
    classPath   *classpath.ClassPath
    classMap    map[string]*Class
}

func (self ClassLoader) loadClass(name string) (*Class) {
    class := self.classMap[name]
    if class != nil {
        return class
    }

    classData, err := self.classPath.ReadClassData(name)
    if err != nil {
        // todo
        panic("class not found:" + name)
    }

    cf, err := classfile.ParseClassFile(classData)
    if err != nil {
        // todo
        panic("failed to parse class file!")
    }

    class = newClass(cf)
    self.classMap[name] = class
    return class
}

func NewClassLoader(cp *classpath.ClassPath) (*ClassLoader) {
    classMap := map[string]*Class{}
    return &ClassLoader{cp, classMap}
}
