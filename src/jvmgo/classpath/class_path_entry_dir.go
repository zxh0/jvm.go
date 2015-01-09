package classpath

import "io/ioutil"

type ClassPathDirEntry struct {
    dir string
}

func (self *ClassPathDirEntry) readClassData(path string) ([]byte, error) {
    fullPath := self.dir + path
    data, err := ioutil.ReadFile(fullPath) 
    if err != nil {
        return nil, err
    }

    return data, nil
}
