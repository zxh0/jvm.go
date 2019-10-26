package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantClass struct {
	name     string
	resolved *Class
}

func newConstantClass(cf *classfile.ClassFile,
	cfc classfile.ConstantClassInfo) *ConstantClass {

	return &ConstantClass{
		name: cf.GetUTF8(cfc.NameIndex),
	}
}

func (ref *ConstantClass) GetClass() *Class {
	if ref.resolved == nil {
		ref.resolve()
	}
	return ref.resolved
}

// todo
func (ref *ConstantClass) resolve() {
	// load class
	ref.resolved = bootLoader.LoadClass(ref.name)
}
