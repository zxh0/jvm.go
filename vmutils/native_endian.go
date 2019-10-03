package vmutils

import (
	"encoding/binary"
	"unsafe"
)

var NativeEndian nativeEndian
var _ binary.ByteOrder = NativeEndian

type nativeEndian struct{}

func (n nativeEndian) Uint16(b []byte) uint16 {
	_ = b[1]
	return *(*uint16)(unsafe.Pointer(&b[0]))
}

func (n nativeEndian) Uint32(b []byte) uint32 {
	_ = b[3]
	return *(*uint32)(unsafe.Pointer(&b[0]))
}

func (n nativeEndian) Uint64(b []byte) uint64 {
	_ = b[7]
	return *(*uint64)(unsafe.Pointer(&b[0]))
}

func (n nativeEndian) PutUint16([]byte, uint16) {
	panic("not implemented")
}

func (n nativeEndian) PutUint32([]byte, uint32) {
	panic("not implemented")
}

func (n nativeEndian) PutUint64([]byte, uint64) {
	panic("not implemented")
}

func (n nativeEndian) String() string {
	panic("not implemented")
}
