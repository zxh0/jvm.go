package class

import (
	"github.com/zxh0/jvm.go/jvmgo/jtype"
	"strings"
)

func GetReturnDescriptor(methodDescriptor string) string {
	start := strings.Index(methodDescriptor, ")") + 1
	return methodDescriptor[start:]
}

func calcArgSlotCount(descriptor string) uint {
	md := parseMethodDescriptor(descriptor)
	slotCount := md.argCount()
	for _, paramType := range md.ParameterTypes() {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
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
	default: // primirive
		return jtype.GetPrimitiveType(descriptor)
	}
}
