package heap

const (
	V = 'V' // void
	Z = 'Z' // boolean
	B = 'B' // byte
	C = 'C' // char
	S = 'S' // short
	I = 'I' // int
	J = 'J' // long
	F = 'F' // float
	D = 'D' // double
)

var PrimitiveTypes = []*PrimitiveType{
	&PrimitiveType{"V", "[V", "void", "java/lang/Void"},
	&PrimitiveType{"Z", "[Z", "boolean", "java/lang/Boolean"},
	&PrimitiveType{"B", "[B", "byte", "java/lang/Byte"},
	&PrimitiveType{"C", "[C", "char", "java/lang/Character"},
	&PrimitiveType{"S", "[S", "short", "java/lang/Short"},
	&PrimitiveType{"I", "[I", "int", "java/lang/Integer"},
	&PrimitiveType{"J", "[J", "long", "java/lang/Long"},
	&PrimitiveType{"F", "[F", "float", "java/lang/Float"},
	&PrimitiveType{"D", "[D", "double", "java/lang/Double"},
}

// type jboolean bool
// type jbyte int8
// type jchar uint16
// type jshort int16
// type jint int32
// type jlong int64
// type jfloat float32
// type jdouble float64

type PrimitiveType struct {
	Descriptor       string
	ArrayClassName   string
	Name             string
	WrapperClassName string
}

func isPrimitiveType(name string) bool {
	for _, primitiveType := range PrimitiveTypes {
		if primitiveType.Name == name {
			return true
		}
	}
	return false
}

func getPrimitiveType(descriptor string) string {
	for _, primitiveType := range PrimitiveTypes {
		if primitiveType.Descriptor == descriptor {
			return primitiveType.Name
		}
	}
	panic("Not primitive type: " + descriptor)
}
