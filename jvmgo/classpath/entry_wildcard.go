package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type WildcardEntry struct {
	compositeEntry CompositeEntry
}

func newWildcardEntry(path string) *WildcardEntry {
	compositeEntry := CompositeEntry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry.addEntry(jarEntry)
		}

		return nil
	}

	dir := path[:len(path)-1]
	filepath.Walk(dir, walkFn)

	return &WildcardEntry{compositeEntry}
}

func (self *WildcardEntry) readClass(className string) (Entry, []byte, error) {
	return self.compositeEntry.readClass(className)
}

func (self *WildcardEntry) String() string {
	return self.compositeEntry.String()
}
