package classpath

import "io/ioutil"

type DirClassPathEntry struct {
	dir string
}

func (self *DirClassPathEntry) readClassData(className string) (ClassPathEntry, []byte, error) {
	fullPath := self.dir + className
	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return self, nil, err
	}

	return self, data, nil
}

func (self *DirClassPathEntry) String() string {
	return self.dir
}
