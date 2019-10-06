package classfile

import (
	"math"

	"github.com/zxh0/jvm.go/vmutils"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantIntegerInfo(reader *ClassReader) int32 {
	return int32(reader.ReadUint32())
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantFloatInfo(reader *ClassReader) float32 {
	return math.Float32frombits(reader.ReadUint32())
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantLongInfo(reader *ClassReader) int64 {
	return int64(reader.ReadUint64())
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantDoubleInfo(reader *ClassReader) float64 {
	return math.Float64frombits(reader.ReadUint64())
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
func readConstantUtf8Info(reader *ClassReader) string {
	length := uint(reader.ReadUint16())
	bytes := reader.ReadBytes(length)
	return vmutils.DecodeMUTF8(bytes)
}
