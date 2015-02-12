package classpath

import (
    "archive/zip"
    "errors"
    "fmt"
    "io/ioutil"
    "jvmgo/jvm/options"
)

type ClassPathJarEntry struct {
    jar     string
    zipRC   *zip.ReadCloser
}

func newClassPathJarEntry(jar string) (*ClassPathJarEntry) {
    return &ClassPathJarEntry{jar, nil}
}

func (self *ClassPathJarEntry) String() string {
    return self.jar
}

func (self *ClassPathJarEntry) readClassData(className string) ([]byte, error) {
    if self.zipRC == nil {
        err := self.openJar()
        if err != nil {
            return nil, err
        }
    }

    classFile := self.findClass(className)
    if classFile == nil {
        return nil, errors.New("class not found!")
    }

    return readClass(classFile)    
}

// todo: close jar
func (self *ClassPathJarEntry) openJar() error {
    r, err := zip.OpenReader(self.jar) // func OpenReader(name string) (*ReadCloser, error)
    if err == nil {
        self.zipRC = r
        if options.VerboseClass {
            fmt.Printf("[Opened %v]\n", self.jar)
        }
    }
    return err
}

func (self *ClassPathJarEntry) findClass(className string) (*zip.File) {
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
