package classpath

import "strings"

type ClassPath struct {
    entries []ClassPathEntry
}

type ClassPathEntry interface {

}

type ClassPathDirEntry struct {
    // todo
}

type ClassPathJarEntry struct {
    // todo
}

func ParseClassPath(cpOption string) (*ClassPath) {
    if cpOption == "" {
        return &ClassPath{}
    }
    //ss := strings.Split(cpOption, ";")

    if strings.Index(cpOption, ";") == -1 {
        //ss.Map()
    }
    // todo
    return nil
}
