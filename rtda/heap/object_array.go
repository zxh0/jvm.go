package heap

import (
	"fmt"
	"reflect"

	"github.com/zxh0/jvm.go/vmutils"
)

const (
	//Array Type  atype
	ATBoolean = 4
	ATChar    = 5
	ATFloat   = 6
	ATDouble  = 7
	ATByte    = 8
	ATShort   = 9
	ATInt     = 10
	ATLong    = 11
)

func NewArray(arrClass *Class, count uint) *Object {
	if arrClass.IsPrimitiveArray() {
		return newPrimitiveArray(arrClass, count)
	} else {
		componentClass := arrClass.GetComponentClass()
		return newRefArray(componentClass, count)
	}
}

func newPrimitiveArray(arrClass *Class, count uint) *Object {
	switch arrClass.Name {
	case "[Z":
		return newObj(arrClass, make([]int8, count), nil)
	case "[B":
		return newObj(arrClass, make([]int8, count), nil)
	case "[C":
		return newObj(arrClass, make([]uint16, count), nil)
	case "[S":
		return newObj(arrClass, make([]int16, count), nil)
	case "[I":
		return newObj(arrClass, make([]int32, count), nil)
	case "[J":
		return newObj(arrClass, make([]int64, count), nil)
	case "[F":
		return newObj(arrClass, make([]float32, count), nil)
	case "[D":
		return newObj(arrClass, make([]float64, count), nil)
	default:
		panic(fmt.Errorf("not primitive array: %v", arrClass))
	}
}

func newRefArray(componentClass *Class, count uint) *Object {
	arrClass := componentClass.getArrayClass()
	components := make([]*Object, count)
	return newObj(arrClass, components, nil)
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	srcArr := reflect.ValueOf(src.Fields)
	dstArr := reflect.ValueOf(dst.Fields)
	reflect.Copy(
		dstArr.Slice(int(dstPos), int(dstPos+length)),
		srcArr.Slice(int(srcPos), int(srcPos+length)),
	)
}

func (obj *Object) ArrayLength() int32 {
	return int32(reflect.ValueOf(obj.Fields).Len())
}

func (obj *Object) GetRefs() []*Object    { return obj.Fields.([]*Object) }
func (obj *Object) GetBooleans() []int8   { return obj.Fields.([]int8) }
func (obj *Object) GetBytes() []int8      { return obj.Fields.([]int8) }
func (obj *Object) GetChars() []uint16    { return obj.Fields.([]uint16) }
func (obj *Object) GetShorts() []int16    { return obj.Fields.([]int16) }
func (obj *Object) GetInts() []int32      { return obj.Fields.([]int32) }
func (obj *Object) GetLongs() []int64     { return obj.Fields.([]int64) }
func (obj *Object) GetFloats() []float32  { return obj.Fields.([]float32) }
func (obj *Object) GetDoubles() []float64 { return obj.Fields.([]float64) }

func (obj *Object) GetGoBytes() []byte {
	s := obj.Fields.([]int8)
	return vmutils.CastInt8sToBytes(s)
}

func (obj *Object) IsArray() bool {
	return obj.Class.IsArray()
}
func (obj *Object) IsPrimitiveArray() bool {
	return obj.Class.IsPrimitiveArray()
}
