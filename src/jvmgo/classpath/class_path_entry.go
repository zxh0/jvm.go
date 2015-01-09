package classpath

import (
    "io/ioutil"
    "strings"
)

type ClassPathEntry interface {
    readClassData(path string) ([]byte, error)
}

type ClassPathDirEntry struct {
    dir string
}

type ClassPathJarEntry struct {
    jar string
}

func (self *ClassPathDirEntry) readClassData(path string) ([]byte, error) {
    fullPath := self.dir + path
    bytes, err := ioutil.ReadFile(fullPath) 
    if err != nil {
        return nil, err
    }

    return bytes, nil
}

func (self *ClassPathJarEntry) readClassData(path string) ([]byte, error) {
    // todo
    return nil, nil
}

func parseClassPathEntry(str string) (ClassPathEntry) {
    if strings.HasSuffix(str, ".jar") {
        return &ClassPathJarEntry{str}
    } else {
        if !strings.HasSuffix(str, "/") {
            str = str + "/"
        }

        return &ClassPathDirEntry{str}
    }
}
