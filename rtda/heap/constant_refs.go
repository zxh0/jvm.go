package heap

import (
	"fmt"
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantClass struct {
	Name     string
	Resolved *Class
}

type ConstantMemberRef struct {
	ClassName     string
	Name          string
	Descriptor    string
	ResolvedClass *Class
}

type ConstantFieldRef struct {
	ConstantMemberRef
	ResolvedField *Field
}

type ConstantMethodRef struct {
	ConstantMemberRef
	IsInterface    bool // is ConstantInterfaceMethodRef ?
	ResolvedMethod *Method
	vslot          int
}

func newConstantClass(cf *classfile.ClassFile, cfc classfile.ConstantClassInfo) *ConstantClass {
	return &ConstantClass{
		Name: cf.GetUTF8(cfc.NameIndex),
	}
}

func newConstantMemberRef(cf *classfile.ClassFile,
	classIdx, nameAndTypeIdx uint16) ConstantMemberRef {

	className := cf.GetClassName(classIdx)
	name, descriptor := cf.GetNameAndType(nameAndTypeIdx)
	return ConstantMemberRef{
		ClassName:  className,
		Name:       name,
		Descriptor: descriptor,
	}
}

func newConstantFieldRef(cf *classfile.ClassFile,
	cfRef classfile.ConstantFieldRefInfo) *ConstantFieldRef {

	return &ConstantFieldRef{
		ConstantMemberRef: newConstantMemberRef(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex),
	}
}

func newConstantMethodRef(cf *classfile.ClassFile,
	cfRef classfile.ConstantMethodRefInfo) *ConstantMethodRef {

	ref := &ConstantMethodRef{
		ConstantMemberRef: newConstantMemberRef(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex),
		IsInterface:       cfRef.IsInterfaceMethod,
		vslot:             -1,
	}
	return ref
}

func (ref *ConstantMemberRef) String() string {
	return fmt.Sprintf("{class:%s, name:%s, descriptor:%s}",
		ref.ClassName, ref.Name, ref.Descriptor) // TODO
}
