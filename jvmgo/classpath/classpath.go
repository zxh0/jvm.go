package classpath

import (
	"path/filepath"
	"strings"

	"github.com/zxh0/jvm.go/jvmgo/options"
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

func (self *ClassPath) parseBootAndExtClassPath() {
	// jre/lib/*
	jreLibPath := filepath.Join(options.AbsJavaHome, "lib", "*")
	self.addEntry(newWildcardEntry(jreLibPath))

	// jre/lib/ext/*
	jreExtPath := filepath.Join(options.AbsJavaHome, "lib", "ext", "*")
	self.addEntry(newWildcardEntry(jreExtPath))
}

func (self *ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.addEntry(newEntry(cpOption))
}

// className: fully/qualified/ClassName
func (self *ClassPath) ReadClass(className string) (Entry, []byte, error) {
	className = className + ".class"
	return self.readClass(className)
}

func (self *ClassPath) String() string {
	userClassPath := self.CompositeEntry.entries[2]
	return userClassPath.String()
}

func IsBootClassPath(entry Entry) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), options.AbsJreLib)
}
