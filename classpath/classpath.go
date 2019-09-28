package classpath

import (
	"path/filepath"
	"strings"

	"github.com/zxh0/jvm.go/options"
)

type ClassPath struct {
	CompositeEntry
}

func Parse(opts options.Options) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(opts.AbsJavaHome)
	cp.parseUserClassPath(opts.Classpath)
	return cp
}

func (cp *ClassPath) parseBootAndExtClassPath(absJavaHome string) {
	// jre/lib/*
	jreLibPath := filepath.Join(absJavaHome, "lib", "*")
	cp.addEntry(newWildcardEntry(jreLibPath))

	// jre/lib/ext/*
	jreExtPath := filepath.Join(absJavaHome, "lib", "ext", "*")
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

func IsBootClassPath(entry Entry, absJreLib string) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), absJreLib)
}
