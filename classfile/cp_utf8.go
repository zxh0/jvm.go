package classfile

import (
	"github.com/zxh0/jvm.go/vmutils"
)

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
