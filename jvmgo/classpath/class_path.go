package classpath

import (
	"errors"
	"github.com/zxh0/jvm.go/jvmgo/jvm/options"
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

func Parse(cpOption string) *ClassPath {
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

func (self *ClassPath) String() string {
	// todo
	return self.compoundEntry.entries[1].String()
}

func IsBootClassPath(entry ClassPathEntry) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), absBootPath)
}
