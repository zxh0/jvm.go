package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantMemberRef struct {
	className  string
	name       string
	descriptor string
}

func (mr *ConstantMemberRef) init(cf *classfile.ClassFile, classIdx, nameAndTypeIdx uint16) {
	mr.className = cf.GetClassNameOf(classIdx)
	mr.name, mr.descriptor = getNameAndType(cf, nameAndTypeIdx)
}
