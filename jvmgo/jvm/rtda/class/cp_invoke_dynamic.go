package class

import (
	"fmt"
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ConstantInvokeDynamic struct {
	name  string
	_type string
	spec  *BootstrapMethodSpecifier
}

type BootstrapMethodSpecifier struct {
	methodRefKind  uint8
	methodRefIndex uint16
	bootstrapArgs  []uint16
}

func newConstantInvokeDynamic(indyInfo *cf.ConstantInvokeDynamicInfo) *ConstantInvokeDynamic {
	name, _type := indyInfo.NameAndType()
	methodRefKind, methodRefIndex, bootstrapArgs := indyInfo.BootstrapMethodInfo()
	return &ConstantInvokeDynamic{
		name:  name,
		_type: _type,
		spec: &BootstrapMethodSpecifier{
			methodRefKind:  methodRefKind,
			methodRefIndex: methodRefIndex,
			bootstrapArgs:  bootstrapArgs,
		},
	}
}

func (self *ConstantInvokeDynamic) String() string {
	return fmt.Sprintf("{name:%v type:%v refKind:%v refIndex:%v args:%v}",
		self.name, self._type, self.spec.methodRefKind, self.spec.methodRefIndex, self.spec.bootstrapArgs)
}

func (self *ConstantInvokeDynamic) Name() string {
	return self.name
}
func (self *ConstantInvokeDynamic) Type() string {
	return self._type
}
func (self *ConstantInvokeDynamic) BootstrapMethodSpecifier() *BootstrapMethodSpecifier {
	return self.spec
}
