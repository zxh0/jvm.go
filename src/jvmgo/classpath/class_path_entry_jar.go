package classpath

import (
    "archive/zip"
    "errors"
)

type ClassPathJarEntry struct {
    jar string
}

func (self *ClassPathJarEntry) readClassData(className string) ([]byte, error) {
    // open jar
    r, err := zip.OpenReader(self.jar) // func OpenReader(name string) (*ReadCloser, error)
    if err != nil {
        return nil, err
    }
    defer r.Close()

    // find class
    className = className + ".class"
    for _, f := range r.File {
        if f.Name == className {
            _, err := f.Open() // func (f *File) Open() (rc io.ReadCloser, err error)
            if err != nil {
                return nil, err
            }
            // read class data
        }
    }

    // todo
    return nil, errors.New("class not found!")
}
