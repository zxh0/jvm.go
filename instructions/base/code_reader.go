package base

import (
	"encoding/binary"

	"github.com/zxh0/jvm.go/vmutils"
)

type CodeReader struct {
	vmutils.BytesReader
}

func NewCodeReader(code []byte) *CodeReader {
	br := vmutils.NewBytesReader(code, binary.BigEndian)
	return &CodeReader{BytesReader: br}
}

func (reader *CodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *CodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *CodeReader) ReadInt32() int32 {
	return int32(reader.ReadUint32())
}

func (reader *CodeReader) ReadInt32s(count int32) []int32 {
	s := make([]int32, count)
	for i := range s {
		s[i] = reader.ReadInt32()
	}
	return s
}

// used by lookupswitch and tableswitch
func (reader *CodeReader) SkipPadding() {
	for reader.Position()%4 != 0 {
		reader.ReadUint8()
	}
}
