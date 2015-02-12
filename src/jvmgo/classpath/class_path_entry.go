package classpath

import (
    //"io/ioutil"
    "strings"
)

type ClassPathEntry interface {
    // className: fully/qualified/ClassName.class
    readClassData(className string) ([]byte, error)
}

func parseClassPathEntry(absPath string) (ClassPathEntry) {
    if strings.HasSuffix(absPath, ".jar") {
        return newClassPathJarEntry(absPath)
    } else {
        if !strings.HasSuffix(absPath, "/") {
            absPath = absPath + "/"
        }

        return &ClassPathDirEntry{absPath}
    }
}
