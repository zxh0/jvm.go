package vmutils

import (
	"os"
	"strings"
)

func IsDir(path string) bool {
	if fileInfo, err := os.Stat(path); err == nil {
		return fileInfo.IsDir()
	}
	return false
}

func IsExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
		//} else if os.IsNotExist(err) {
		//	return false
	} else {
		return false
	}
}

func IsZipFile(name string) bool {
	return strings.HasSuffix(name, ".zip") ||
		strings.HasSuffix(name, ".ZIP")
}

func IsJarFile(name string) bool {
	return strings.HasSuffix(name, ".jar") ||
		strings.HasSuffix(name, ".JAR")
}

func IsJModFile(name string) bool {
	return strings.HasSuffix(name, ".jmod")
}
