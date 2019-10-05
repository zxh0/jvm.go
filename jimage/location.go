package jimage

import (
	"fmt"

	"github.com/zxh0/jvm.go/vmutils"
)

const (
	AttributeEnd          = 0
	AttributeModule       = 1
	AttributeParent       = 2
	AttributeBase         = 3
	AttributeExtension    = 4
	AttributeOffset       = 5
	AttributeCompressed   = 6
	AttributeUncompressed = 7
	AttributeCount        = 8
)

type Location struct {
	attributes []uint64
	image      *Image
}

func (location Location) GetModuleOffset() uint64 {
	return location.getAttribute(AttributeModule)
}
func (location Location) GetModule() string {
	return location.getAttributeString(AttributeModule)
}

func (location Location) GetBaseOffset() uint64 {
	return location.getAttribute(AttributeBase)
}
func (location Location) GetBase() string {
	return location.getAttributeString(AttributeBase)
}

func (location Location) GetParentOffset() uint64 {
	return location.getAttribute(AttributeParent)
}
func (location Location) GetParent() string {
	return location.getAttributeString(AttributeParent)
}

func (location Location) GetExtensionOffset() uint64 {
	return location.getAttribute(AttributeExtension)
}
func (location Location) GetExtension() string {
	return location.getAttributeString(AttributeExtension)
}

func (location Location) GetContentOffset() uint64 {
	return location.getAttribute(AttributeOffset)
}

func (location Location) GetCompressedSize() uint64 {
	return location.getAttribute(AttributeCompressed)
}

func (location Location) GetUncompressedSize() uint64 {
	return location.getAttribute(AttributeUncompressed)
}

func (location Location) getAttributeString(kind int) string {
	offset := location.getAttribute(kind)
	return location.image.getString(int(offset))
}

func (location Location) getAttribute(kind int) uint64 {
	if kind < AttributeEnd || AttributeCount <= kind {
		panic(fmt.Errorf("invalid jimage attribute kind: %d", kind))
	}
	return location.attributes[kind]
}

func (location Location) getFullName(modulesPrefix bool) string {
	var fullName vmutils.StringBuilder

	if location.GetModuleOffset() != 0 {
		if modulesPrefix {
			fullName.Append("/modules")
		}
		fullName.Append("/", location.GetModule(), "/")
	}

	if location.GetParentOffset() != 0 {
		fullName.Append(location.GetParent(), "/")
	}

	fullName.Append(location.GetBase())

	if location.GetExtensionOffset() != 0 {
		fullName.Append(".", location.GetExtension())
	}

	return fullName.String()
}

// static long[] decompress(ByteBuffer bytes)
// https://github.com/unofficial-openjdk/openjdk/blob/jdk/jdk/src/java.base/share/classes/jdk/internal/jimage/ImageLocation.java
func decompress(bytes []byte) []uint64 {
	attributes := make([]uint64, AttributeCount)

	nBytes := len(bytes)
	for i := 0; i < nBytes; i++ {
		data := bytes[i]
		kind := data >> 3

		if kind == AttributeEnd {
			break
		}

		if kind < AttributeEnd || AttributeCount <= kind {
			panic(fmt.Errorf("invalid jimage attribute kind: %d", kind))
		}

		length := int((data & 0x7) + 1)
		var value uint64 = 0

		for j := 0; j < length; j++ {
			value <<= 8

			if i == nBytes-1 {
				panic(fmt.Errorf("missing jimage attribute data"))
			}

			i++
			value |= uint64(bytes[i])
		}

		attributes[kind] = value
	}

	return attributes
}
