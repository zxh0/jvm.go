package heap

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ConstantClass struct {
	name  string
	class *Class
}

func newConstantClass(classInfo *cf.ConstantClassInfo) *ConstantClass {
	return &ConstantClass{
		name: classInfo.Name(),
	}
}

func (self *ConstantClass) Class() *Class {
	if self.class == nil {
		self.resolve()
	}
	return self.class
}

// todo
func (self *ConstantClass) resolve() {
	// load class
	self.class = bootLoader.LoadClass(self.name)
}
