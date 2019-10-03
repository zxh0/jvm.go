package jimage

import (
	"github.com/zxh0/jvm.go/vmutils"
)

func NewImageReader(data []byte) vmutils.BytesReader {
	return vmutils.NewBytesReader(data, vmutils.NativeEndian)
}

func ReadHeader(reader vmutils.BytesReader) Header {
	magic := reader.ReadUint32()
	version := reader.ReadUint32()
	flags := reader.ReadUint32()
	resourceCount := reader.ReadUint32()
	tableLength := reader.ReadUint32()
	locationsSize := reader.ReadUint32()
	stringsSize := reader.ReadUint32()

	return Header{
		Magic:         magic,
		MajorVersion:  uint16(version >> 16),
		MinorVersion:  uint16(version & 0xFFFF),
		Flags:         flags,
		ResourceCount: resourceCount,
		TableLength:   tableLength,
		LocationsSize: locationsSize,
		StringsSize:   stringsSize,
	}
}
