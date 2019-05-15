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

func newConstantMemberref(refInfo *cf.ConstantMemberrefInfo) Constant {
	switch refInfo.Tag {
	case cf.CONSTANT_Fieldref:
		ref := &ConstantFieldref{}
		ref.copy(refInfo)
		return ref
	case cf.CONSTANT_Methodref:
		ref := &ConstantMethodref{vslot: -1}
		ref.copy(refInfo)
		ref.argSlotCount = calcArgSlotCount(ref.descriptor)
		return ref
	case cf.CONSTANT_InterfaceMethodref:
		ref := &ConstantInterfaceMethodref{}
		ref.copy(refInfo)
		ref.argSlotCount = calcArgSlotCount(ref.descriptor)
		return ref
	default:
		panic("unreachable!")
	}
}
