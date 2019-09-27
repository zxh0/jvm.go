package base

import "encoding/binary"

type CodeReader struct {
	code []byte // bytecodes
	pc   int
}

func NewCodeReader(code []byte) *CodeReader {
	return &CodeReader{code: code}
}

func (reader *CodeReader) PC() int {
	return reader.pc
}

func (reader *CodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}
func (reader *CodeReader) ReadUint8() uint8 {
	i := reader.code[reader.pc]
	reader.pc++
	return i
}

func (reader *CodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}
func (reader *CodeReader) ReadUint16() uint16 {
	i := binary.BigEndian.Uint16(reader.code[reader.pc:])
	reader.pc += 2
	return i
}

func (reader *CodeReader) ReadInt32() int32 {
	i := binary.BigEndian.Uint32(reader.code[reader.pc:])
	reader.pc += 4
	return int32(i)
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
	for reader.pc%4 != 0 {
		reader.ReadUint8()
	}
}
