package vmutils

import (
	"encoding/binary"
	"unsafe"
)

var NativeEndian binary.ByteOrder

// https://stackoverflow.com/questions/51332658/any-better-way-to-check-endianness-in-go
func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		NativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		NativeEndian = binary.BigEndian
	default:
		panic("Could not determine native endianness.")
	}
}

func IsBigEndian() bool {
	return NativeEndian == binary.BigEndian
}
