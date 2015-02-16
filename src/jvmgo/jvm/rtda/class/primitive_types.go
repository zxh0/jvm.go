package class

import "jvmgo/util"

// java primitive types
var primitiveTypes = map[string]string{
	// name     descriptor
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"char":    "C",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"float":   "F",
	"double":  "D",
}

func isPrimitiveType(name string) bool {
	_, ok := primitiveTypes[name]
	return ok
}

func getPrimitiveType(descriptor string) string {
	for name, desc := range primitiveTypes {
		if desc == descriptor {
			return name
		}
	}
	util.Panicf("Not primitive type: %v", descriptor)
	return ""
}
