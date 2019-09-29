package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantMemberRef struct {
	className  string
	name       string
	descriptor string
}

func newConstantMemberRef(cf *classfile.ClassFile, refInfo classfile.ConstantMemberRefInfo) Constant {
	switch refInfo.Tag {
	case classfile.ConstantFieldRef:
		ref := &ConstantFieldRef{}
		ref.copy(cf, refInfo)
		return ref
	case classfile.ConstantMethodRef:
		ref := &ConstantMethodRef{vslot: -1}
		ref.copy(cf, refInfo)
		ref.ArgSlotCount = calcArgSlotCount(ref.descriptor)
		return ref
	case classfile.ConstantInterfaceMethodRef:
		ref := &ConstantInterfaceMethodRef{}
		ref.copy(cf, refInfo)
		ref.ArgSlotCount = calcArgSlotCount(ref.descriptor)
		return ref
	default:
		panic("unreachable!")
	}
}

func (mr *ConstantMemberRef) copy(cf *classfile.ClassFile, refInfo classfile.ConstantMemberRefInfo) {
	mr.className = cf.GetClassNameOf(refInfo.ClassIndex)
	mr.name, mr.descriptor = getNameAndType(cf, refInfo.NameAndTypeIndex)
}
