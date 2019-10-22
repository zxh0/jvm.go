package classpath

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/zxh0/jvm.go/vmutils"
)

type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, error)
	String() string
}

func parsePath(path string) []Entry {
	switch {
	case strings.IndexByte(path, os.PathListSeparator) >= 0:
		return splitPath(path)
	case strings.HasSuffix(path, "*"):
		return spreadWildcardEntry(path)
	case vmutils.IsJarFile(path) || vmutils.IsZipFile(path):
		return []Entry{newZipEntry(path)}
	default:
		return []Entry{newDirEntry(path)}
	}
}

func splitPath(pathList string) []Entry {
	list := make([]Entry, 0, 4)

	for _, path := range strings.Split(pathList, string(os.PathListSeparator)) {
		list = append(list, parsePath(path)...)
	}

	return list
}

func spreadWildcardEntry(path string) []Entry {
	baseDir := path[:len(path)-1] // remove *
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		panic(err) // TODO
	}

	list := make([]Entry, 0, 4)
	for _, file := range files {
		if vmutils.IsJarFile(file.Name()) {
			filename := filepath.Join(baseDir, file.Name())
			list = append(list, newZipEntry(filename))
		}
	}

	return list
}
