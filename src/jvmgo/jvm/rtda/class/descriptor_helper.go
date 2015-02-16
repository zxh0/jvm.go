package class

import (
	"jvmgo/util"
	"strings"
)

func calcArgCount(descriptor string) uint {
	return parseMethodDescriptor(descriptor).argCount()
}

func getComponentDescriptor(descriptor string) string {
	if descriptor[0] != '[' {
		util.Panicf("Not array: %v", descriptor)
		return ""
	}
	return descriptor[1:]
}

func getClassName(descriptor string) string {
	switch descriptor[0] {
	case '[':
		return descriptor // array
	case 'L':
		return descriptor[1 : len(descriptor)-1] // object
	default:
		return getPrimitiveType(descriptor) // primirive types
	}
}

func GetReturnDescriptor(methodDescriptor string) string {
	start := strings.Index(methodDescriptor, ")") + 1
	return methodDescriptor[start:]
}
