package classpath

import (
    "archive/zip"
    "fmt"
)

type ClassPathJarEntry struct {
    jar string
}

func (self *ClassPathJarEntry) readClassData(path string) ([]byte, error) {
    r, err := zip.OpenReader(self.jar)
    if err != nil {
        return nil, err
    }
    defer r.Close()

    for _, f := range r.File {
        if f.Name == path {
            // todo
            fmt.Printf("jar f: %s \n", f.Name)
        }
    }

    // todo
    return nil, nil
}
