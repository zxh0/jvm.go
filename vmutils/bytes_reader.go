package vmutils

import (
	"encoding/binary"
)

type BytesReader struct {
	byteOrder binary.ByteOrder
	data      []byte
	position  int
}

func NewBytesReader(data []byte, byteOrder binary.ByteOrder) BytesReader {
	return BytesReader{
		byteOrder: byteOrder,
		data:      data,
		position:  0,
	}
}

func (reader *BytesReader) Position() int {
	return reader.position
}

func (reader *BytesReader) ReadUint8() uint8 {
	i := reader.data[reader.position]
	reader.position++
	return i
}

func (reader *BytesReader) ReadUint16() uint16 {
	i := reader.byteOrder.Uint16(reader.data[reader.position:])
	reader.position += 2
	return i
}

func (reader *BytesReader) ReadUint32() uint32 {
	i := reader.byteOrder.Uint32(reader.data[reader.position:])
	reader.position += 4
	return i
}

func (reader *BytesReader) ReadUint64() uint64 {
	i := reader.byteOrder.Uint64(reader.data[reader.position:])
	reader.position += 8
	return i
}

func (reader *BytesReader) ReadBytes(n int) []byte {
	bytes := reader.data[reader.position : reader.position+n]
	reader.position += n
	return bytes
}
