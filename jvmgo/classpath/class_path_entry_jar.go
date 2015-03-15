package classpath

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"jvmgo/jvm/options"
)

type JarClassPathEntry struct {
	jar   string
	zipRC *zip.ReadCloser
}

func newJarClassPathEntry(jar string) *JarClassPathEntry {
	return &JarClassPathEntry{jar, nil}
}

func (self *JarClassPathEntry) readClassData(className string) (ClassPathEntry, []byte, error) {
	if self.zipRC == nil {
		err := self.openJar()
		if err != nil {
			return self, nil, err
		}
	}

	classFile := self.findClass(className)
	if classFile == nil {
		return self, nil, classNotFoundErr
	}

	data, err := readClass(classFile)
	return self, data, err
}

// todo: close jar
func (self *JarClassPathEntry) openJar() error {
	r, err := zip.OpenReader(self.jar) // func OpenReader(name string) (*ReadCloser, error)
	if err == nil {
		self.zipRC = r
		if options.VerboseClass {
			fmt.Printf("[Opened %v]\n", self.jar)
		}
	}
	return err
}

func (self *JarClassPathEntry) findClass(className string) *zip.File {
	for _, f := range self.zipRC.File {
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

func (self *JarClassPathEntry) String() string {
	return self.jar
}
