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
	compositeEntry CompositeEntry
}

func Parse(pathList string) *ClassPath {
	if pathList == "" {
		pathList = "."
	}

	// jre/lib/*
	jreLibPath := filepath.Join(options.AbsJavaHome, "lib", "*")

	return &ClassPath{
		CompositeEntry{
			[]Entry{
				newWildcardEntry(jreLibPath), // boot classpath
				newCompositeEntry(pathList),
			},
		},
	}
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClass(className string) (Entry, []byte, error) {
	className = className + ".class"
	return self.compositeEntry.readClass(className)
}

func (self *ClassPath) String() string {
	// todo
	return self.compositeEntry.entries[1].String()
}

func IsBootClassPath(entry Entry) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), options.AbsJreLib)
}
