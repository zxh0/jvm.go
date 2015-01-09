package classpath

import "strings"

type ClassPath struct {
    entries []ClassPathEntry
}

type ClassPathEntry interface {

}

type ClassPathDirEntry struct {
    dir string
}

type ClassPathJarEntry struct {
    jar string
}

func ParseClassPath(cpOption string) (*ClassPath) {
    if cpOption == "" {
        return &ClassPath{}
    }

    cpOptionSplitted := strings.Split(cpOption, ";")
    cpEntries := make([]ClassPathEntry, len(cpOptionSplitted))

    for idx, str := range cpOptionSplitted {
        cpEntries[idx] = parseClassPathEntry(str)
    }

    return &ClassPath{cpEntries}
}

func parseClassPathEntry(str string) (ClassPathEntry) {
    if strings.HasSuffix(str, ".jar") {
        return &ClassPathJarEntry{str}
    } else {
        return &ClassPathDirEntry{str}
    }
}
