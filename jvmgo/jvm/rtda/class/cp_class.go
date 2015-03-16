package class

import cf "github.com/zxh0/jvm.go/jvmgo/classfile"

type ConstantClass struct {
	name  string
	cp    *ConstantPool
	class *Class
}

func newConstantClass(cp *ConstantPool, classInfo *cf.ConstantClassInfo) *ConstantClass {
	return &ConstantClass{
		name: classInfo.Name(),
		cp:   cp,
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
