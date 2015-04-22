package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type AsteriskEntry struct {
	compoundEntry CompoundEntry
}

func newAsteriskEntry(path string) *AsteriskEntry {
	compoundEntry := CompoundEntry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newJarEntry(path)
			compoundEntry.addEntry(jarEntry)
		}

		return nil
	}

	dir := path[:len(path)-1]
	filepath.Walk(dir, walkFn)

	return &AsteriskEntry{compoundEntry}
}

func (self *AsteriskEntry) readClassData(className string) (Entry, []byte, error) {
	return self.compoundEntry.readClassData(className)
}

func (self *AsteriskEntry) String() string {
	return self.compoundEntry.String()
}
