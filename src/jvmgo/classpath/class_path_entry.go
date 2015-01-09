package classpath

import "strings"

type ClassPathEntry interface {
    ReadClassData(path string) ([]byte)
}

type ClassPathDirEntry struct {
    dir string
}

type ClassPathJarEntry struct {
    jar string
}

func (self *ClassPathDirEntry) ReadClassData(path string) ([]byte) {
    // todo
    return nil
}

func (self *ClassPathJarEntry) ReadClassData(path string) ([]byte) {
    // todo
    return nil
}

func parseClassPathEntry(str string) (ClassPathEntry) {
    if strings.HasSuffix(str, ".jar") {
        return &ClassPathJarEntry{str}
    } else {
        return &ClassPathDirEntry{str}
    }
}

