package class

import (
	"fmt"
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

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

func (self *ConstantInvokeDynamic) String() string {
	return fmt.Sprintf("{name:%v type:%v refKind:%v refIndex:%v args:%v}",
		self.name, self._type, self.refKind, self.refIndex, self.args)
}

func (self *ConstantInvokeDynamic) Name() string {
	return self.name
}
func (self *ConstantInvokeDynamic) Type() string {
	return self._type
}
func (self *ConstantInvokeDynamic) RefKind() uint8 {
	return self.refKind
}
func (self *ConstantInvokeDynamic) RefIndex() uint16 {
	return self.refIndex
}
func (self *ConstantInvokeDynamic) Args() []uint16 {
	return self.args
}
