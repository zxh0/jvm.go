package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantInvokeDynamic struct {
	name               string
	_type              string
	bootstrapMethodRef uint16 // method handle
	bootstrapArguments []uint16
	cp                 ConstantPool
}

func newConstantInvokeDynamic(cf *classfile.ClassFile, cp ConstantPool, indyInfo classfile.ConstantInvokeDynamicInfo) *ConstantInvokeDynamic {
	name, _type := cf.GetNameAndType(indyInfo.NameAndTypeIndex)
	bm := cf.GetBootstrapMethods()[indyInfo.BootstrapMethodAttrIndex]
	return &ConstantInvokeDynamic{
		name:               name,
		_type:              _type,
		bootstrapMethodRef: bm.BootstrapMethodRef,
		bootstrapArguments: bm.BootstrapArguments,
		cp:                 cp,
	}
}

func (indy *ConstantInvokeDynamic) MethodHandle() {
	kMH := indy.cp.GetConstant(uint(indy.bootstrapMethodRef)).(*ConstantMethodHandle)
	println(kMH)
}
