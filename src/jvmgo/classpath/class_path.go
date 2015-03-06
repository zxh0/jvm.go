package classpath

import (
	"errors"
)

var classNotFoundErr = errors.New("class not found!")

type ClassPath struct {
	compoundEntry CompoundClassPathEntry
}

func ParseClassPath(cpOption string) *ClassPath {
	if cpOption == "" {
		return &ClassPath{}
	}

	compoundEntry := parseCompoundClassPathEntry(cpOption)
	return &ClassPath{compoundEntry}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClassData(className string) (ClassPathEntry, []byte, error) {
	className = className + ".class"
	return self.compoundEntry.readClassData(className)
}
