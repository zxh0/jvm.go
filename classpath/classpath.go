package classpath

import (
	"path/filepath"
	"strings"

	"github.com/zxh0/jvm.go/vm"
)

type ClassPath struct {
	entries []Entry
}

func Parse(opts *vm.Options) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(opts.AbsJavaHome)
	cp.parseUserClassPath(opts.ClassPath)
	return cp
}

func (cp *ClassPath) parseBootAndExtClassPath(absJavaHome string) {
	// jre/lib/*
	jreLibPath := filepath.Join(absJavaHome, "lib", "*")
	cp.entries = append(cp.entries, spreadWildcardEntry(jreLibPath)...)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(absJavaHome, "lib", "ext", "*")
	cp.entries = append(cp.entries, spreadWildcardEntry(jreExtPath)...)
}

func (cp *ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.entries = append(cp.entries, parsePath(cpOption)...)
}

// className: fully/qualified/ClassName
func (cp *ClassPath) ReadClass(className string) (Entry, []byte) {
	className = className + ".class"
	for _, entry := range cp.entries {
		if data, err := entry.readClass(className); err == nil {
			return entry, data
		}
	}
	return nil, nil
}

func IsBootClassPath(entry Entry, absJreLib string) bool {
	if entry == nil {
		// todo
		return true
	}

	return strings.HasPrefix(entry.String(), absJreLib)
}
