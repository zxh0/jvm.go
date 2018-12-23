package heap

import (
	"strings"
)

func GetReturnDescriptor(methodDescriptor string) string {
	start := strings.Index(methodDescriptor, ")") + 1
	return methodDescriptor[start:]
}

func calcArgSlotCount(descriptor string) uint {
	return parseMethodDescriptor(descriptor).argSlotCount()
}

// [XXX -> [XXX
// LXXX; -> XXX
// I -> int ...
func getClassName(descriptor string) string {
	switch descriptor[0] {
	case '[': // array
		return descriptor
	case 'L': // object
		return descriptor[1 : len(descriptor)-1]
	default: // primitive
		return getPrimitiveType(descriptor)
	}
}
