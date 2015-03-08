package classpath

import (
	"errors"
	"jvmgo/jvm/options"
	"path/filepath"
	"strings"
)

var (
	absBootPath      = filepath.Join(options.AbsJavaHome, "lib") // jre/lib
	classNotFoundErr = errors.New("class not found!")
)

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
				parseCompoundClassPathEntry(filepath.Join(absBootPath, "*")), // jre/lib/*
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

func IsBootClassPath(entry ClassPathEntry) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), absBootPath)
}
