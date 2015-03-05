package classpath

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

type ClassPath struct {
	entries []ClassPathEntry
}

func ParseClassPath(cpOption string) *ClassPath {
	if cpOption == "" {
		return &ClassPath{}
	}

	cpOptionSplitted := strings.Split(cpOption, pathListSeparator)
	cpEntries := make([]ClassPathEntry, len(cpOptionSplitted))

	for i, p := range cpOptionSplitted {
		absPath, err := filepath.Abs(p)
		if err == nil {
			cpEntries[i] = parseClassPathEntry(absPath)
		} else {
			// todo
		}
	}

	return &ClassPath{cpEntries}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClassData(className string) (ClassPathEntry, []byte, error) {
	className = className + ".class"
	for _, entry := range self.entries {
		entry, data, err := entry.readClassData(className)
		if err == nil {
			return entry, data, nil
		}
	}

	// todo
	err := errors.New("class not found!")
	return nil, nil, err
}
