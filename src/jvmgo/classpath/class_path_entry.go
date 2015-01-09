package classpath

import (
    //"io/ioutil"
    "strings"
)

type ClassPathEntry interface {
    // className: path/to/ClassFile
    readClassData(className string) ([]byte, error)
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
