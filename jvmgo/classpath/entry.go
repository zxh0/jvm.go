package classpath

import (
	"strings"
)

type ClassPathEntry interface {
	// className: fully/qualified/ClassName.class
	readClassData(className string) (ClassPathEntry, []byte, error)
	String() string
}

func parseClassPathEntry(absPath string) ClassPathEntry {
	if strings.HasSuffix(absPath, "*") {
		return newAsteriskClassPathEntry(absPath)
	}

	if strings.HasSuffix(absPath, ".jar") {
		return newJarClassPathEntry(absPath)
	}

	return newDirClassPathEntry(absPath)
}
