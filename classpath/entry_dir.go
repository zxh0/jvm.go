package classpath

import (
	"github.com/zxh0/jvm.go/vmutils"
)

type DirEntry struct {
	dir *vmutils.Dir
}

func newDirEntry(path string) *DirEntry {
	if dir, err := vmutils.NewDir(path); err != nil {
		panic(err) // TODO
	} else {
		return &DirEntry{dir: dir}
	}
}

func (entry *DirEntry) readClass(className string) ([]byte, error) {
	return entry.dir.ReadFile(className)
}

func (entry *DirEntry) String() string {
	return entry.dir.AbsPath()
}
