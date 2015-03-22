package classpath

import (
	"errors"
	"github.com/zxh0/jvm.go/jvmgo/jvm/options"
	"path/filepath"
	"strings"
)

var (
	classNotFoundErr = errors.New("class not found!")
)

type ClassPath struct {
	compoundEntry CompoundClassPathEntry
}

func Parse(pathList string) *ClassPath {
	if pathList == "" {
		pathList = "."
	}

	// jre/lib/*
	jreLibPath := filepath.Join(options.AbsJavaHome, "lib", "*")

	return &ClassPath{
		CompoundClassPathEntry{
			[]ClassPathEntry{
				newCompoundClassPathEntry(jreLibPath), // boot classpath
				newCompoundClassPathEntry(pathList),
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

	return strings.HasPrefix(entry.String(), options.AbsJreLib)
}
