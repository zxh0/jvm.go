package classpath

import "io/ioutil"

type ClassPathDirEntry struct {
    dir string
}

func (self *ClassPathDirEntry) readClassData(className string) ([]byte, error) {
    fullPath := self.dir + className
    data, err := ioutil.ReadFile(fullPath) 
    if err != nil {
        return nil, err
    }

    return data, nil
}
