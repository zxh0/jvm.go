package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantClass struct {
	name  string
	class *Class
}

func newConstantClass(cf *classfile.ClassFile, cfc classfile.ConstantClassInfo) *ConstantClass {
	return &ConstantClass{
		name: cf.GetUTF8(cfc.NameIndex),
	}
}

func (cr *ConstantClass) GetClass() *Class {
	if cr.class == nil {
		cr.resolve()
	}
	return cr.class
}

// todo
func (cr *ConstantClass) resolve() {
	// load class
	cr.class = bootLoader.LoadClass(cr.name)
}
