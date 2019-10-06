package jimage

import (
	"fmt"
	"math"
	"sort"
)

type Image struct {
	Header
	redirect  []int32
	offsets   []uint32
	locations []byte
	strings   []byte
	fullData  []byte
}

func (image Image) GetEntryNames() []string {
	names := make([]string, 0, 100)
	for _, offset := range image.offsets {
		if offset > 0 {
			location := Location{
				attributes: image.getAttributes(offset),
				image:      &image,
			}
			name := location.getFullName(false)
			names = append(names, name)
		}
	}
	sort.Sort(sort.StringSlice(names))
	return names
}

func (image Image) FindLocation(name string) *Location {
	count := int32(image.TableLength)
	index := image.redirect[hashCode(name, HashMultiplier)%count]

	if index < 0 {
		// index is twos complement of location attributes index.
		index = -index - 1
	} else if index > 0 {
		// index is hash seed needed to compute location attributes index.
		index = hashCode(name, index) % count
	} else {
		// No entry.
		return nil
	}

	attributes := image.getAttributes(image.offsets[index])

	// TODO
	//if !ImageLocation.verify(module, name, attributes, stringsReader) {
	//	return nil
	//}
	return &Location{
		attributes: attributes,
		image:      &image,
	}
}

func (image Image) GetResource(name string) []byte {
	if location := image.FindLocation(name); location == nil {
		return nil
	} else {
		return image.getResource0(location)
	}
}

func (image Image) getResource0(loc *Location) []byte {
	offset := loc.GetContentOffset() + uint64(image.GetIndexSize())
	compressedSize := loc.GetCompressedSize()
	uncompressedSize := loc.GetUncompressedSize()

	if compressedSize < 0 || math.MaxInt32 < compressedSize {
		panic(fmt.Errorf("bad compressed size: %d", compressedSize))
	}
	if uncompressedSize < 0 || math.MaxInt32 < uncompressedSize {
		panic(fmt.Errorf("bad uncompressed size: %d", uncompressedSize))
	}
	if compressedSize != 0 {
		panic("resource is compressed") // TODO
	}

	return image.fullData[offset : offset+uncompressedSize]
}

func (image Image) getAttributes(offset uint32) []uint64 {
	//if offset < 0 || offset >= len(image.locations) {
	//	panic("invalid offset")
	//}
	return decompress(image.locations[offset:])
}

func (image Image) getString(offset int) string {
	bytes := getStringBytes(image.strings, offset)
	return string(bytes) // TODO
}

func (image Image) DebugListStrings() {
	offset := 0
	for offset < len(image.strings) {
		raw := getStringBytes(image.strings, offset)
		offset += len(raw) + 1
		fmt.Println(string(raw))
	}
}
