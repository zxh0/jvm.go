package heap

import (
	"fmt"
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
		return NewRefArray(componentClass, count)
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

func NewRefArray(componentClass *Class, count uint) *Object {
	arrClass := componentClass.arrayClass()
	components := make([]*Object, count)
	return newObj(arrClass, components, nil)
}

// todo rename
func NewRefArray2(componentClass *Class, components []*Object) *Object {
	arrClass := componentClass.arrayClass()
	return newObj(arrClass, components, nil)
}

func ArrayLength(arr *Object) int32 {
	switch x := arr.fields.(type) {
	case []int8:
		return int32(len(x))
	case []int16:
		return int32(len(x))
	case []int32:
		return int32(len(x))
	case []int64:
		return int32(len(x))
	case []uint16:
		return int32(len(x))
	case []float32:
		return int32(len(x))
	case []float64:
		return int32(len(x))
	case []*Object:
		return int32(len(x))
	default:
		panic(fmt.Errorf("not array: %v", arr))
	}
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.fields.(type) {
	case []int8:
		_src := src.fields.([]int8)[srcPos : srcPos+length]
		_dst := dst.fields.([]int8)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.fields.([]int16)[srcPos : srcPos+length]
		_dst := dst.fields.([]int16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.fields.([]int32)[srcPos : srcPos+length]
		_dst := dst.fields.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int64:
		_src := src.fields.([]int64)[srcPos : srcPos+length]
		_dst := dst.fields.([]int64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.fields.([]uint16)[srcPos : srcPos+length]
		_dst := dst.fields.([]uint16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.fields.([]float32)[srcPos : srcPos+length]
		_dst := dst.fields.([]float32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.fields.([]float64)[srcPos : srcPos+length]
		_dst := dst.fields.([]float64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.fields.([]*Object)[srcPos : srcPos+length]
		_dst := dst.fields.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
	default:
		panic(fmt.Errorf("not array: %v", src))
	}
}
