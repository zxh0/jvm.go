package class

import cf "github.com/zxh0/jvm.go/jvmgo/classfile"

type ConstantInvokeDynamic struct {
	name  string
	_type string
}

func newConstantInvokeDynamic(indyInfo *cf.ConstantInvokeDynamicInfo) *ConstantInvokeDynamic {
	name, _type := indyInfo.NameAndType()
	return &ConstantInvokeDynamic{
		name:  name,
		_type: _type,
	}
}
