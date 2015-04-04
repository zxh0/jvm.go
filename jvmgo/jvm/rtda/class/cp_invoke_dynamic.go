package class

import cf "github.com/zxh0/jvm.go/jvmgo/classfile"

type ConstantInvokeDynamic struct {
	name     string
	_type    string
	refKind  uint8
	refIndex uint16
	args     []uint16
}

func newConstantInvokeDynamic(indyInfo *cf.ConstantInvokeDynamicInfo) *ConstantInvokeDynamic {
	name, _type := indyInfo.NameAndType()
	refKind, refIndex, args := indyInfo.BootstrapMethodInfo()
	return &ConstantInvokeDynamic{
		name:     name,
		_type:    _type,
		refKind:  refKind,
		refIndex: refIndex,
		args:     args,
	}
}
