package classpath

import (
	"errors"
	"jvmgo/jvm/options"
)

var classNotFoundErr = errors.New("class not found!")

type ClassPath struct {
	compoundEntry CompoundClassPathEntry
}

func ParseClassPath(cpOption string) *ClassPath {
	if cpOption == "" {
		cpOption = "."
	}

	return &ClassPath{
		CompoundClassPathEntry{
			[]ClassPathEntry{
				parseCompoundClassPathEntry(options.JavaHome + "lib/*"),
				parseCompoundClassPathEntry(cpOption),
			},
		},
	}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClassData(className string) (ClassPathEntry, []byte, error) {
	className = className + ".class"
	return self.compoundEntry.readClassData(className)
}
