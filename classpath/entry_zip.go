package classpath

import (
	"github.com/zxh0/jvm.go/vmutils"
)

type ZipEntry struct {
	zipFile *vmutils.ZipFile
}

func newZipEntry(path string) *ZipEntry {
	if zipFile, err := vmutils.NewZipFile(path); err != nil {
		panic(err) // TODO
	} else {
		return &ZipEntry{zipFile: zipFile}
	}
}

func (entry *ZipEntry) readClass(className string) ([]byte, error) {
	// TODO: close ZipFile
	if !entry.zipFile.IsOpen() {
		if err := entry.zipFile.Open(); err != nil {
			return nil, err
		}
	}

	return entry.zipFile.ReadFile(className)
}

func (entry *ZipEntry) String() string {
	return entry.zipFile.AbsPath()
}
