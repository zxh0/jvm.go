package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantMemberRef struct {
	class      *Class
	className  string
	name       string
	descriptor string
}

func newConstantMemberRef(class *Class, cf *classfile.ClassFile,
	classIdx, nameAndTypeIdx uint16) ConstantMemberRef {

	className := cf.GetClassName(classIdx)
	name, descriptor := cf.GetNameAndType(nameAndTypeIdx)
	return ConstantMemberRef{
		class:      class,
		className:  className,
		name:       name,
		descriptor: descriptor,
	}
}

func (ref *ConstantMemberRef) getBootLoader() *ClassLoader {
	return ref.class.bootLoader
}
