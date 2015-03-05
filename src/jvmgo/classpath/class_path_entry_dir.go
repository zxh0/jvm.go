package classpath

import "io/ioutil"

type DirClassPathEntry struct {
	dir string
}

func (self *DirClassPathEntry) readClassData(className string) ([]byte, error) {
	fullPath := self.dir + className
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (self *DirClassPathEntry) String() string {
	return self.dir
}
