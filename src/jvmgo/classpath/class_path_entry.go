package classpath

import "strings"

type ClassPathEntry interface {

}

type ClassPathDirEntry struct {
    dir string
}

type ClassPathJarEntry struct {
    jar string
}

func parseClassPathEntry(str string) (ClassPathEntry) {
    if strings.HasSuffix(str, ".jar") {
        return &ClassPathJarEntry{str}
    } else {
        return &ClassPathDirEntry{str}
    }
}
