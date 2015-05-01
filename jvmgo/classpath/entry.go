package classpath

import (
	"strings"
)

type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) (Entry, []byte, error)
	String() string
}

func parseEntry(absPath string) Entry {
	if strings.HasSuffix(absPath, "*") {
		return newWildcardEntry(absPath)
	}

	if strings.HasSuffix(absPath, ".jar") {
		return newJarEntry(absPath)
	}

	return newDirEntry(absPath)
}
