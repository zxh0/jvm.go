package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (entry *DirEntry) readClass(className string) (Entry, []byte, error) {
	fileName := filepath.Join(entry.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return entry, nil, err
	}
	return entry, data, nil
}

func (entry *DirEntry) String() string {
	return entry.absDir
}
