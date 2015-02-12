package classpath

import (
    "errors"
    "os"
    "strings"
)

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

type ClassPath struct {
    entries []ClassPathEntry
}

func ParseClassPath(cpOption string) (*ClassPath) {
    if cpOption == "" {
        return &ClassPath{}
    }

    cpOptionSplitted := strings.Split(cpOption, pathListSeparator)
    cpEntries := make([]ClassPathEntry, len(cpOptionSplitted))

    for i, str := range cpOptionSplitted {
        cpEntries[i] = parseClassPathEntry(str)
    }

    return &ClassPath{cpEntries}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClassData(className string) ([]byte, error) {
    className = className + ".class"
    for _, entry := range self.entries {
        data, err := entry.readClassData(className)
        if err == nil {
            return data, nil
        }
    }

    // todo
    err := errors.New("class not found!")
    return nil, err
}
