package classpath

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

var classNotFoundErr = errors.New("class not found!")

type ClassPath struct {
	compoundEntry CompoundClassPathEntry
}

func ParseClassPath(cpOption string) *ClassPath {
	if cpOption == "" {
		return &ClassPath{}
	}

	compoundEntry := CompoundClassPathEntry{}
	for _, path := range strings.Split(cpOption, pathListSeparator) {
		if absPath, err := filepath.Abs(path); err == nil {
			entry := parseClassPathEntry(absPath)
			compoundEntry.addEntry(entry)
		} else {
			// todo
		}
	}

	return &ClassPath{compoundEntry}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClassData(className string) (ClassPathEntry, []byte, error) {
	className = className + ".class"
	return self.compoundEntry.readClassData(className)
}
