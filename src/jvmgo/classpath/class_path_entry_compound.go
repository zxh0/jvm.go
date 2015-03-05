package classpath

// import (
// 	"errors"
// )

// type CompoundClassPathEntry struct {
// 	entries []ClassPathEntry
// }

// // className: fully/qualified/ClassName
// func (self *CompoundClassPathEntry) ReadClassData(className string) (ClassPathEntry, []byte, error) {
// 	className = className + ".class"
// 	for _, entry := range self.entries {
// 		data, err := entry.readClassData(className)
// 		if err == nil {
// 			return entry, data, nil
// 		}
// 	}

// 	// todo
// 	err := errors.New("class not found!")
// 	return nil, nil, err
// }
