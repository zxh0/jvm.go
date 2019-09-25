package classpath

import (
	"path/filepath"
	"strings"

	"github.com/zxh0/jvm.go/options"
)

type ClassPath struct {
	CompositeEntry
}

func Parse(cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath()
	cp.parseUserClassPath(cpOption)
	return cp
}

func (cp *ClassPath) parseBootAndExtClassPath() {
	// jre/lib/*
	jreLibPath := filepath.Join(options.AbsJavaHome, "lib", "*")
	cp.addEntry(newWildcardEntry(jreLibPath))

	// jre/lib/ext/*
	jreExtPath := filepath.Join(options.AbsJavaHome, "lib", "ext", "*")
	cp.addEntry(newWildcardEntry(jreExtPath))
}

func (cp *ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.addEntry(newEntry(cpOption))
}

// className: fully/qualified/ClassName
func (cp *ClassPath) ReadClass(className string) (Entry, []byte, error) {
	className = className + ".class"
	return cp.readClass(className)
}

func (cp *ClassPath) String() string {
	userClassPath := cp.CompositeEntry.entries[2]
	return userClassPath.String()
}

func IsBootClassPath(entry Entry) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), options.AbsJreLib)
}
