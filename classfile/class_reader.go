package classfile

import (
	"encoding/binary"

	"github.com/zxh0/jvm.go/vmutils"
)

type ClassReader struct {
	vmutils.BytesReader
	cf *ClassFile
}

func newClassReader(data []byte) ClassReader {
	br := vmutils.NewBytesReader(data, binary.BigEndian)
	return ClassReader{BytesReader: br}
}

func (reader *ClassReader) ReadUint16s() []uint16 {
	n := reader.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = reader.ReadUint16()
	}
	return s
}
