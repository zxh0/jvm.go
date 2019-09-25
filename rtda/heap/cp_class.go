package heap

import (
	cf "github.com/zxh0/jvm.go/classfile"
)

type ConstantClass struct {
	name  string
	class *Class
}

func newConstantClass(classInfo cf.ConstantClassInfo) *ConstantClass {
	return &ConstantClass{
		name: classInfo.Name(),
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
