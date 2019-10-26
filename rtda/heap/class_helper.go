package heap

// type jboolean bool
// type jbyte int8
// type jchar uint16
// type jshort int16
// type jint int32
// type jlong int64
// type jfloat float32
// type jdouble float64

var primitiveTypes = []PrimitiveType{
	{"V", "[V", "void", "java/lang/Void"},
	{"Z", "[Z", "boolean", "java/lang/Boolean"},
	{"B", "[B", "byte", "java/lang/Byte"},
	{"C", "[C", "char", "java/lang/Character"},
	{"S", "[S", "short", "java/lang/Short"},
	{"I", "[I", "int", "java/lang/Integer"},
	{"J", "[J", "long", "java/lang/Long"},
	{"F", "[F", "float", "java/lang/Float"},
	{"D", "[D", "double", "java/lang/Double"},
}

type PrimitiveType struct {
	Descriptor       string
	ArrayClassName   string
	Name             string
	WrapperClassName string // not used
}

func isPrimitiveType(name string) bool {
	for _, primitiveType := range primitiveTypes {
		if primitiveType.Name == name {
			return true
		}
	}
	return false
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
		return getPrimitiveClassName(descriptor)
	}
}

func getPrimitiveClassName(descriptor string) string {
	for _, primitiveType := range primitiveTypes {
		if primitiveType.Descriptor == descriptor {
			return primitiveType.Name
		}
	}
	panic("Not primitive type: " + descriptor)
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	if className[0] == '[' {
		// array
		return "[" + className
	}
	for _, primitiveType := range primitiveTypes {
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

// TODO
func GetPrimitiveDescriptor(className string) string {
	switch className {
	case "java/lang/Boolean":
		return "Z"
	case "java/lang/Byte":
		return "B"
	case "java/lang/Character":
		return "C"
	case "java/lang/Short":
		return "S"
	case "java/lang/Integer":
		return "I"
	case "java/lang/Long":
		return "J"
	case "java/lang/Float":
		return "F"
	case "java/lang/Double":
		return "D"
	default:
		return ""
	}
}
