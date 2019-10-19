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
		return _newPrimitiveArray(arrClass, count)
	} else {
		componentClass := arrClass.ComponentClass()
		return NewRefArrayN(componentClass, count)
	}
}

func _newPrimitiveArray(arrClass *Class, count uint) *Object {
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
	case "[L":
		return newObj(arrClass, make([]int64, count), nil)
	case "[F":
		return newObj(arrClass, make([]float32, count), nil)
	case "[D":
		return newObj(arrClass, make([]float64, count), nil)
	default:
		panic(fmt.Errorf("not primitive array: %v", arrClass))
	}
}

func NewPrimitiveArray(atype uint8, count uint) *Object {
	switch atype {
	case ATBoolean:
		return newObj(bootLoader.getClass("[Z"), make([]int8, count), nil)
	case ATByte:
		return newObj(bootLoader.getClass("[B"), make([]int8, count), nil)
	case ATChar:
		return newObj(bootLoader.getClass("[C"), make([]uint16, count), nil)
	case ATShort:
		return newObj(bootLoader.getClass("[S"), make([]int16, count), nil)
	case ATInt:
		return newObj(bootLoader.getClass("[I"), make([]int32, count), nil)
	case ATLong:
		return newObj(bootLoader.getClass("[J"), make([]int64, count), nil)
	case ATFloat:
		return newObj(bootLoader.getClass("[F"), make([]float32, count), nil)
	case ATDouble:
		return newObj(bootLoader.getClass("[D"), make([]float64, count), nil)
	default:
		panic(fmt.Errorf("invalid atype: %v", atype))
	}
}

func NewByteArray(bytes []int8) *Object {
	return newObj(bootLoader.getClass("[B"), bytes, nil)
}
func NewCharArray(chars []uint16) *Object {
	return newObj(bootLoader.getClass("[C"), chars, nil)
}
func NewRefArray(componentClass *Class, components []*Object) *Object {
	arrClass := componentClass.arrayClass()
	return newObj(arrClass, components, nil)
}
func NewRefArrayN(componentClass *Class, count uint) *Object {
	arrClass := componentClass.arrayClass()
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
