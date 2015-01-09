package classpath

import (
    "archive/zip"
    "errors"
    "io/ioutil"
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
    for _, f := range r.File {
        if f.Name == className {
            rc, err := f.Open() // func (f *File) Open() (rc io.ReadCloser, err error)
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
    }

    // todo
    return nil, errors.New("class not found!")
}
