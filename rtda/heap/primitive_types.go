package heap

var PrimitiveTypes = []PrimitiveType{
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
