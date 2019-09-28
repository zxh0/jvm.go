package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantClass struct {
	name  string
	class *Class
}

func newConstantClass(cf *classfile.ClassFile, classInfo classfile.ConstantClassInfo) *ConstantClass {
	return &ConstantClass{
		name: cf.GetUTF8(classInfo.NameIndex),
	}
}

func (cc *ConstantClass) Class() *Class {
	if cc.class == nil {
		cc.resolve()
	}
	return cc.class
}

// todo
func (cc *ConstantClass) resolve() {
	// load class
	cc.class = bootLoader.LoadClass(cc.name)
}
