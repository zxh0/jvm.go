package heap

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

const (
	//Array Type  atype
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
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
	switch arrClass.Name() {
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
		jutil.Panicf("Not primitive array: %v!", arrClass)
		return nil
	}
}

func NewPrimitiveArray(atype uint8, count uint) *Object {
	switch atype {
	case AT_BOOLEAN:
		return newObj(bootLoader.getClass("[Z"), make([]int8, count), nil)
	case AT_BYTE:
		return newObj(bootLoader.getClass("[B"), make([]int8, count), nil)
	case AT_CHAR:
		return newObj(bootLoader.getClass("[C"), make([]uint16, count), nil)
	case AT_SHORT:
		return newObj(bootLoader.getClass("[S"), make([]int16, count), nil)
	case AT_INT:
		return newObj(bootLoader.getClass("[I"), make([]int32, count), nil)
	case AT_LONG:
		return newObj(bootLoader.getClass("[J"), make([]int64, count), nil)
	case AT_FLOAT:
		return newObj(bootLoader.getClass("[F"), make([]float32, count), nil)
	case AT_DOUBLE:
		return newObj(bootLoader.getClass("[D"), make([]float64, count), nil)
	default:
		jutil.Panicf("BAD atype: %v!", atype)
		return nil
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
	switch arr.fields.(type) {
	case []int8:
		return int32(len(arr.fields.([]int8)))
	case []int16:
		return int32(len(arr.fields.([]int16)))
	case []int32:
		return int32(len(arr.fields.([]int32)))
	case []int64:
		return int32(len(arr.fields.([]int64)))
	case []uint16:
		return int32(len(arr.fields.([]uint16)))
	case []float32:
		return int32(len(arr.fields.([]float32)))
	case []float64:
		return int32(len(arr.fields.([]float64)))
	case []*Object:
		return int32(len(arr.fields.([]*Object)))
	default:
		jutil.Panicf("Not array: %v!", arr)
		return -1
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
		jutil.Panicf("Not array: %v!", src)
	}
}
