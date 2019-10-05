package jimage

import (
	"fmt"
	"sort"
)

type Image struct {
	Header
	redirect  []int32
	offsets   []uint32
	locations []byte
	strings   []byte
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

func (image Image) FindLocation2(module, name string) *Location {
	count := int32(image.TableLength)
	index := image.redirect[hashCode2(module, name, HashMultiplier)%count]

	if index < 0 {
		// index is twos complement of location attributes index.
		index = -index - 1
	} else if index > 0 {
		// index is hash seed needed to compute location attributes index.
		index = hashCode2(module, name, index) % count
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
