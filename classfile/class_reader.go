package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func newClassReader(data []byte) ClassReader {
	return ClassReader{data}
}

// u1
func (reader *ClassReader) readUint8() uint8 {
	val := reader.data[0]
	reader.data = reader.data[1:]
	return val
}

// u2
func (reader *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(reader.data)
	reader.data = reader.data[2:]
	return val
}

// u4
func (reader *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(reader.data)
	reader.data = reader.data[4:]
	return val
}

func (reader *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(reader.data)
	reader.data = reader.data[8:]
	return val
}

func (reader *ClassReader) readUint16s() []uint16 {
	n := reader.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = reader.readUint16()
	}
	return s
}

func (reader *ClassReader) readBytes(length uint32) []byte {
	bytes := reader.data[:length]
	reader.data = reader.data[length:]
	return bytes
}
