package classpath

import "strings"

type ClassPath struct {
    entries []ClassPathEntry
}

func (self *ClassPath) ReadClassData(path string) ([]byte) {
    for _, entry := range self.entries {

    }

    // todo
    return nil
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
