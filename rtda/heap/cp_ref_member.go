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

	className := cf.GetClassNameOf(classIdx)
	name, descriptor := getNameAndType(cf, nameAndTypeIdx)
	return ConstantMemberRef{
		class:      class,
		className:  className,
		name:       name,
		descriptor: descriptor,
	}
}
