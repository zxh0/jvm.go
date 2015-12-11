package heap

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

func DotToSlash(name string) string {
	return jutil.ReplaceAll(name, ".", "/")
}
func SlashToDot(name string) string {
	return jutil.ReplaceAll(name, "/", ".")
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	if className[0] == '[' {
		// array
		return "[" + className
	}
	for _, primitiveType := range PrimitiveTypes {
		if primitiveType.Name == className {
			// primitive
			return primitiveType.ArrayClassName
		}
	}
	// object
	return "[L" + className + ";"
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	if className[0] == '[' {
		descriptor := className[1:]
		return getClassName(descriptor)
	}
	panic("Not array: " + className)
}
