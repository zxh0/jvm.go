package classpath

import (
    //"io/ioutil"
    "strings"
)

type ClassPathEntry interface {
    // className: fully/qualified/ClassName.class
    readClassData(className string) ([]byte, error)
}

func parseClassPathEntry(str string) (ClassPathEntry) {
    if strings.HasSuffix(str, ".jar") {
        return newClassPathJarEntry(str)
    } else {
        if !strings.HasSuffix(str, "/") {
            str = str + "/"
        }

        return &ClassPathDirEntry{str}
    }
}
