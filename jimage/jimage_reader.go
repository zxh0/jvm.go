package jimage

import (
	"github.com/zxh0/jvm.go/vmutils"
)

func ReadHeader(data []byte) Header {
	reader := vmutils.NewBytesReader(data, vmutils.NativeEndian)
	return readHeader(reader)
}

func readHeader(reader vmutils.BytesReader) Header {
	magic := reader.ReadUint32()
	version := reader.ReadUint32()
	flags := reader.ReadUint32()
	resourceCount := reader.ReadUint32()
	tableLength := reader.ReadUint32()
	locationsSize := reader.ReadUint32()
	stringsSize := reader.ReadUint32()

	header := Header{
		Magic:         magic,
		MajorVersion:  uint16(version >> 16),
		MinorVersion:  uint16(version & 0xFFFF),
		Flags:         flags,
		ResourceCount: resourceCount,
		TableLength:   tableLength,
		LocationsSize: locationsSize,
		StringsSize:   stringsSize,
	}
	checkHeader(header)
	return header
}

func checkHeader(header Header) { // TODO
	if header.Magic != Magic {
		panic("not an image file")
	}
	if header.MajorVersion != MajorVersion ||
		header.MinorVersion != MinorVersion {

		panic("the image file is not the correct version")
	}
}

func ReadImage(data []byte) Image {
	reader := vmutils.NewBytesReader(data, vmutils.NativeEndian)
	header := readHeader(reader)

	redirect := data[header.getRedirectOffset() : header.getRedirectOffset()+header.GetRedirectSize()]
	offsets := data[header.getOffsetsOffset() : header.getOffsetsOffset()+header.GetOffsetsSize()/4] // FIXME
	locations := data[header.getLocationsOffset() : header.getLocationsOffset()+header.LocationsSize]
	strings := data[header.getStringsOffset() : header.getStringsOffset()+header.StringsSize]

	return Image{
		Header:    header,
		redirect:  vmutils.CastBytesToInt32s(redirect),
		offsets:   vmutils.CastBytesToUint32s(offsets),
		locations: locations,
		strings:   strings,
		fullData:  data,
	}
}
