package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantMemberRef struct {
	className  string
	name       string
	descriptor string
}

func newConstantFieldRef(cf *classfile.ClassFile, refInfo classfile.ConstantFieldRefInfo) *ConstantFieldRef {
	ref := &ConstantFieldRef{}
	ref.init(cf, refInfo.ClassIndex, refInfo.NameAndTypeIndex)
	return ref
}

func newConstantMethodRef(cf *classfile.ClassFile, refInfo classfile.ConstantMethodRefInfo) *ConstantMethodRef {
	ref := &ConstantMethodRef{vslot: -1}
	ref.init(cf, refInfo.ClassIndex, refInfo.NameAndTypeIndex)
	ref.ArgSlotCount = calcArgSlotCount(ref.descriptor)
	return ref
}

func newConstantInterfaceMethodRef(cf *classfile.ClassFile, refInfo classfile.ConstantInterfaceMethodRefInfo) *ConstantInterfaceMethodRef {
	ref := &ConstantInterfaceMethodRef{}
	ref.init(cf, refInfo.ClassIndex, refInfo.NameAndTypeIndex)
	ref.ArgSlotCount = calcArgSlotCount(ref.descriptor)
	return ref
}

func (mr *ConstantMemberRef) init(cf *classfile.ClassFile, classIdx, nameAndTypeIdx uint16) {
	mr.className = cf.GetClassNameOf(classIdx)
	mr.name, mr.descriptor = getNameAndType(cf, nameAndTypeIdx)
}
