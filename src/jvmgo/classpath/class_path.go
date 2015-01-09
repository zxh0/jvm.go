package classpath

import (
    "strings"
    "errors"
)

type ClassPath struct {
    entries []ClassPathEntry
}

func (self *ClassPath) ReadClassData(path string) ([]byte, error) {
    for _, entry := range self.entries {
        data, err := entry.readClassData(path)
        if err == nil {
            return data, nil
        }
    }

    // todo
    err := errors.New("class not found!")
    return nil, err
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
