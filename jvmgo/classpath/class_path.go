package classpath

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/zxh0/jvm.go/jvmgo/jvm/options"
)

var (
	classNotFoundErr = errors.New("class not found!")
)

type ClassPath struct {
	compoundEntry CompoundEntry
}

func Parse(pathList string) *ClassPath {
	if pathList == "" {
		pathList = "."
	}

	// jre/lib/*
	jreLibPath := filepath.Join(options.AbsJavaHome, "lib", "*")

	return &ClassPath{
		CompoundEntry{
			[]Entry{
				newCompoundEntry(jreLibPath), // boot classpath
				newCompoundEntry(pathList),
			},
		},
	}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClassData(className string) (Entry, []byte, error) {
	className = className + ".class"
	return self.compoundEntry.readClassData(className)
}

func (self *ClassPath) String() string {
	// todo
	return self.compoundEntry.entries[1].String()
}

func IsBootClassPath(entry Entry) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), options.AbsJreLib)
}
