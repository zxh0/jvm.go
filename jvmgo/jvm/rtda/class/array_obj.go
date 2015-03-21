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

func (self *Obj) Booleans() []int8 {
	return self.fields.([]int8)
}

func (self *Obj) Bytes() []int8 {
	return self.fields.([]int8)
}

func (self *Obj) Chars() []uint16 {
	return self.fields.([]uint16)
}

func (self *Obj) Shorts() []int16 {
	return self.fields.([]int16)
}

func (self *Obj) Ints() []int32 {
	return self.fields.([]int32)
}

func (self *Obj) Longs() []int64 {
	return self.fields.([]int64)
}

func (self *Obj) Floats() []float32 {
	return self.fields.([]float32)
}

func (self *Obj) Doubles() []float64 {
	return self.fields.([]float64)
}

func (self *Obj) GoBytes() []byte {
	s := self.fields.([]int8)
	return util.CastInt8sToUint8s(s)
}
