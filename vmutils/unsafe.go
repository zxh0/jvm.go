package vmutils

import (
	"reflect"
	"unsafe"
)

// []int8 <-> []byte
func CastInt8sToBytes(s []int8) []byte {
	ptr := unsafe.Pointer(&s)
	return *((*[]byte)(ptr))
}
func CastBytesToInt8s(s []byte) []int8 {
	ptr := unsafe.Pointer(&s)
	return *((*[]int8)(ptr))
}

// []int8 <-> []uint16
func CastInt8sToUint16s(s []int8) []uint16 {
	ptr := unsafe.Pointer(&s)
	(*reflect.SliceHeader)(ptr).Len /= 2
	return *((*[]uint16)(ptr))
}
func CastUint16sToInt8s(s []uint16) []int8 {
	ptr := unsafe.Pointer(&s)
	(*reflect.SliceHeader)(ptr).Len *= 2
	return *((*[]int8)(ptr))
}

func CastBytesToUint32s(s []byte) []uint32 {
	ptr := unsafe.Pointer(&s)
	(*reflect.SliceHeader)(ptr).Len /= 4
	return *((*[]uint32)(ptr))
}

func CastBytesToInt32s(s []byte) []int32 {
	ptr := unsafe.Pointer(&s)
	(*reflect.SliceHeader)(ptr).Len /= 4
	return *((*[]int32)(ptr))
}
