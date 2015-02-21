package class

import (
	"jvmgo/util"
	"strings"
)

func calcArgCount(descriptor string) uint {
	return parseMethodDescriptor(descriptor).argCount()
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	if className[0] == '[' {
		// array
		return "[" + className
	}
	descriptor, isPrimitive := primitiveTypes[className]
	if isPrimitive {
		// primitive
		return "[" + descriptor
	}
	// object
	return "[L" + className + ";"
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	cd := getComponentDescriptor(className)
	return getClassName(cd)
}

// [XXX -> XXX
func getComponentDescriptor(descriptor string) string {
	if descriptor[0] != '[' {
		util.Panicf("Not array: %v", descriptor)
		return ""
	}
	return descriptor[1:]
}

// [XXX -> [XXX
// LXXX; -> XXX
// I -> int ...
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
