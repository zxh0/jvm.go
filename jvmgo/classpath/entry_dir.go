package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	dir string
}

func newDirEntry(dir string) *DirEntry {
	return &DirEntry{dir}
}

func (self *DirEntry) readClass(className string) (Entry, []byte, error) {
	fileName := filepath.Join(self.dir, className)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return self, nil, err
	}

	return self, data, nil
}

func (self *DirEntry) String() string {
	return self.dir
}
