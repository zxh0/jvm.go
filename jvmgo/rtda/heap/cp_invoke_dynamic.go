package heap

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ConstantInvokeDynamic struct {
	name               string
	_type              string
	bootstrapMethodRef uint16 // method handle
	bootstrapArguments []uint16
	cp                 *ConstantPool
}

func newConstantInvokeDynamic(cp *ConstantPool, indyInfo *cf.ConstantInvokeDynamicInfo) *ConstantInvokeDynamic {
	name, _type := indyInfo.NameAndType()
	bootstrapMethodRef, bootstrapArguments := indyInfo.BootstrapMethodInfo()
	return &ConstantInvokeDynamic{
		name:               name,
		_type:              _type,
		bootstrapMethodRef: bootstrapMethodRef,
		bootstrapArguments: bootstrapArguments,
		cp:                 cp,
	}
}

func (self *ConstantInvokeDynamic) MethodHandle() {
	println(2222)
	println(self.bootstrapMethodRef)
	kMH := self.cp.GetConstant(uint(self.bootstrapMethodRef)).(*ConstantMethodHandle)
	println(kMH.referenceKind)
	println(kMH.referenceIndex)
	println(kMH)
}
