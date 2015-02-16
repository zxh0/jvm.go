package bigendian

import (
	"encoding/binary"
	"math"
)

var _bigEndian = binary.BigEndian

func PutInt8(s []byte, val int8) {
	s[0] = uint8(val)
}
func Int8(s []byte) int8 {
	return int8(s[0])
}

func PutUint16(s []byte, val uint16) {
	_bigEndian.PutUint16(s, val)
}
func Uint16(s []byte) uint16 {
	return _bigEndian.Uint16(s)
}

func PutInt16(s []byte, val int16) {
	_bigEndian.PutUint16(s, uint16(val))
}
func Int16(s []byte) int16 {
	return int16(_bigEndian.Uint16(s))
}

func PutInt32(s []byte, val int32) {
	_bigEndian.PutUint32(s, uint32(val))
}
func Int32(s []byte) int32 {
	return int32(_bigEndian.Uint32(s))
}

func PutInt64(s []byte, val int64) {
	_bigEndian.PutUint64(s, uint64(val))
}
func Int64(s []byte) int64 {
	return int64(_bigEndian.Uint64(s))
}

func PutFloat32(s []byte, val float32) {
	_bigEndian.PutUint32(s, math.Float32bits(val))
}
func Float32(s []byte) float32 {
	return math.Float32frombits(_bigEndian.Uint32(s))
}

func PutFloat64(s []byte, val float64) {
	_bigEndian.PutUint64(s, math.Float64bits(val))
}
func Float64(s []byte) float64 {
	return math.Float64frombits(_bigEndian.Uint64(s))
}
