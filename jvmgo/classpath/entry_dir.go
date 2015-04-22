package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirClassPathEntry struct {
	dir string
}

func newDirClassPathEntry(dir string) *DirClassPathEntry {
	return &DirClassPathEntry{dir}
}

func (self *DirClassPathEntry) readClassData(className string) (ClassPathEntry, []byte, error) {
	fileName := filepath.Join(self.dir, className)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return self, nil, err
	}

	return self, data, nil
}

func (self *DirClassPathEntry) String() string {
	return self.dir
}
