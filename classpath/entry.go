package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// :(linux/unix) or ;(windows)
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, error)
	String() string
}

func parsePath(path string) []Entry {
	switch {
	case strings.Contains(path, pathListSeparator):
		return splitPath(path)
	case strings.HasSuffix(path, "*"):
		return spreadWildcardEntry(path)
	case strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP"):
		return []Entry{newZipEntry(path)}
	default:
		return []Entry{newDirEntry(path)}
	}
}

func splitPath(pathList string) []Entry {
	list := make([]Entry, 0, 4)

	for _, path := range strings.Split(pathList, pathListSeparator) {
		list = append(list, parsePath(path)...)
	}

	return list
}

func spreadWildcardEntry(path string) []Entry {
	baseDir := path[:len(path)-1] // remove *
	list := make([]Entry, 0, 4)

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") {

			list = append(list, newZipEntry(path))
		}
		return nil
	}

	_ = filepath.Walk(baseDir, walkFn)
	return list
}
