package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type AsteriskClassPathEntry struct {
	compoundEntry CompoundClassPathEntry
}

func newAsteriskClassPathEntry(path string) *AsteriskClassPathEntry {
	compoundEntry := CompoundClassPathEntry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newJarClassPathEntry(path)
			compoundEntry.addEntry(jarEntry)
		}

		return nil
	}

	dir := path[:len(path)-1]
	filepath.Walk(dir, walkFn)

	return &AsteriskClassPathEntry{compoundEntry}
}

func (self *AsteriskClassPathEntry) readClassData(className string) (ClassPathEntry, []byte, error) {
	return self.compoundEntry.readClassData(className)
}

func (self *AsteriskClassPathEntry) String() string {
	return self.compoundEntry.String()
}
