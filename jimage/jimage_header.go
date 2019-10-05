package jimage

const (
	Magic        = 0xCAFEDADA
	MajorVersion = 1
	MinorVersion = 0
	HeaderSlots  = 7
)

/*
header    --+
redirect    |
offsets     |> index
locations   |
strings   --+

*/

// https://github.com/unofficial-openjdk/openjdk/blob/jdk/jdk/src/java.base/share/classes/jdk/internal/jimage/ImageHeader.java
type Header struct {
	Magic         uint32
	MajorVersion  uint16
	MinorVersion  uint16
	Flags         uint32
	ResourceCount uint32
	TableLength   uint32
	LocationsSize uint32
	StringsSize   uint32
}

func getHeaderSize() uint32 {
	return HeaderSlots * 4
}

func (header Header) GetRedirectSize() uint32 {
	return header.TableLength * 4
}

func (header Header) GetOffsetsSize() uint32 {
	return header.TableLength * 4
}

func (header Header) GetIndexSize() uint32 {
	return getHeaderSize() +
		header.GetRedirectSize() +
		header.GetOffsetsSize() +
		header.LocationsSize +
		header.StringsSize
}

func (header Header) getRedirectOffset() uint32 {
	return getHeaderSize()
}

func (header Header) getOffsetsOffset() uint32 {
	return header.getRedirectOffset() +
		header.GetRedirectSize()
}

func (header Header) getLocationsOffset() uint32 {
	return header.getOffsetsOffset() +
		header.GetOffsetsSize()
}

func (header Header) getStringsOffset() uint32 {
	return header.getLocationsOffset() +
		header.LocationsSize
}
