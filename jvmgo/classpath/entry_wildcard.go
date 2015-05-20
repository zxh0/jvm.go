package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type WildcardEntry struct {
	compoundEntry CompoundEntry
}

func newWildcardEntry(path string) *WildcardEntry {
	compoundEntry := CompoundEntry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compoundEntry.addEntry(jarEntry)
		}

		return nil
	}

	dir := path[:len(path)-1]
	filepath.Walk(dir, walkFn)

	return &WildcardEntry{compoundEntry}
}

func (self *WildcardEntry) readClass(className string) (Entry, []byte, error) {
	return self.compoundEntry.readClass(className)
}

func (self *WildcardEntry) String() string {
	return self.compoundEntry.String()
}
