package heap

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

func (self *Object) IsArray() bool {
	return self.class.IsArray()
}
func (self *Object) IsPrimitiveArray() bool {
	return self.class.IsPrimitiveArray()
}

func (self *Object) Refs() []*Object {
	return self.fields.([]*Object)
}

func (self *Object) Booleans() []int8 {
	return self.fields.([]int8)
}

func (self *Object) Bytes() []int8 {
	return self.fields.([]int8)
}

func (self *Object) Chars() []uint16 {
	return self.fields.([]uint16)
}

func (self *Object) Shorts() []int16 {
	return self.fields.([]int16)
}

func (self *Object) Ints() []int32 {
	return self.fields.([]int32)
}

func (self *Object) Longs() []int64 {
	return self.fields.([]int64)
}

func (self *Object) Floats() []float32 {
	return self.fields.([]float32)
}

func (self *Object) Doubles() []float64 {
	return self.fields.([]float64)
}

func (self *Object) GoBytes() []byte {
	s := self.fields.([]int8)
	return jutil.CastInt8sToUint8s(s)
}
