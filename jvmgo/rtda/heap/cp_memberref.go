package heap

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ConstantMemberref struct {
	className  string
	name       string
	descriptor string
}

func (self *ConstantMemberref) copy(refInfo *cf.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}
