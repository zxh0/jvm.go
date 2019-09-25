package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absZip string
	zipRC  *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absZip, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absZip, nil}
}

func (entry *ZipEntry) readClass(className string) (Entry, []byte, error) {
	if entry.zipRC == nil {
		err := entry.openJar()
		if err != nil {
			return entry, nil, err
		}
	}

	classFile := entry.findClass(className)
	if classFile == nil {
		return entry, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return entry, data, err
}

// todo: close zip
func (entry *ZipEntry) openJar() error {
	r, err := zip.OpenReader(entry.absZip) // func OpenReader(name string) (*ReadCloser, error)
	if err == nil {
		entry.zipRC = r
	}
	return err
}

func (entry *ZipEntry) findClass(className string) *zip.File {
	for _, f := range entry.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open() // func (f *File) Open() (rc io.ReadCloser, err error)
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc) // func ReadAll(r io.Reader) ([]byte, error)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (entry *ZipEntry) String() string {
	return entry.absZip
}
