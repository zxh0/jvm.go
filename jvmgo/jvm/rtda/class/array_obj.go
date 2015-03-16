package class

import (
	"github.com/zxh0/jvm.go/jvmgo/util"
)

func (self *Obj) IsArray() bool {
	return self.class.IsArray()
}
func (self *Obj) IsPrimitiveArray() bool {
	return self.class.IsPrimitiveArray()
}

func (self *Obj) Refs() []*Obj {
	return self.fields.([]*Obj)
}

func (self *Obj) GoBytes() []byte {
	s := self.fields.([]int8)
	return util.CastInt8sToUint8s(s)
}
