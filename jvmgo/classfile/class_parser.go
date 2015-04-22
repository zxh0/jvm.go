package classfile

import (
	"fmt"
)

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := newClassReader(classData)
	cf = &ClassFile{}
	cf.read(cr)
	return
}
