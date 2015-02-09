package bigendian

import (
    "encoding/binary"
    "math"
)

var _bigEndian = binary.BigEndian

func EncodeInt16(s []byte, val int16) {
    _bigEndian.PutUint16(s, uint16(val))
}
func DecodeInt16(s []byte) int16 {
    return int16(_bigEndian.Uint16(s))
}

func EncodeInt32(s []byte, val int32) {
    _bigEndian.PutUint32(s, uint32(val))
}
func DecodeInt32(s []byte) int32 {
    return int32(_bigEndian.Uint32(s))
}

func EncodeInt64(s []byte, val int64) {
    _bigEndian.PutUint64(s, uint64(val))
}
func DecodeInt64(s []byte) int64 {
    return int64(_bigEndian.Uint64(s))
}

func EncodeFloat32(s []byte, val float32) {
    _bigEndian.PutUint32(s, math.Float32bits(val))
}
func DecodeFloat32(s []byte) float32 {
    return math.Float32frombits(_bigEndian.Uint32(s))
}

func EncodeFloat64(s []byte, val float64) {
    _bigEndian.PutUint64(s, math.Float64bits(val))
}
func DecodeFloat64(s []byte) float64 {
    return math.Float64frombits(_bigEndian.Uint64(s))
}
