package jutil

import (
	"unsafe"
)

func CastInt8sToUint8s(jBytes []int8) (goBytes []byte) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}

func CastUint8sToInt8s(goBytes []byte) (jBytes []int8) {
	ptr := unsafe.Pointer(&goBytes)
	jBytes = *((*[]int8)(ptr))
	return
}
