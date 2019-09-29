package heap

import (
	"github.com/zxh0/jvm.go/vmutils"
)

func (obj *Object) Refs() []*Object    { return obj.fields.([]*Object) }
func (obj *Object) Booleans() []int8   { return obj.fields.([]int8) }
func (obj *Object) Bytes() []int8      { return obj.fields.([]int8) }
func (obj *Object) Chars() []uint16    { return obj.fields.([]uint16) }
func (obj *Object) Shorts() []int16    { return obj.fields.([]int16) }
func (obj *Object) Ints() []int32      { return obj.fields.([]int32) }
func (obj *Object) Longs() []int64     { return obj.fields.([]int64) }
func (obj *Object) Floats() []float32  { return obj.fields.([]float32) }
func (obj *Object) Doubles() []float64 { return obj.fields.([]float64) }

func (obj *Object) IsArray() bool {
	return obj.class.IsArray()
}
func (obj *Object) IsPrimitiveArray() bool {
	return obj.class.IsPrimitiveArray()
}

func (obj *Object) GoBytes() []byte {
	s := obj.fields.([]int8)
	return vmutils.CastInt8sToUint8s(s)
}
