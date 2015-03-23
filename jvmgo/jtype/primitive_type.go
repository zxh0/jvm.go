package jtypes

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

var (
	Jvoid    = &PrimitiveType{"V", "", "void", "java/lang/Void"}
	Jboolean = &PrimitiveType{"Z", "[Z", "boolean", "java/lang/Boolean"}
	Jbyte    = &PrimitiveType{"B", "[B", "byte", "java/lang/Byte"}
	Jchar    = &PrimitiveType{"C", "[C", "char", "java/lang/Character"}
	Jshort   = &PrimitiveType{"S", "[S", "short", "java/lang/Short"}
	Jint     = &PrimitiveType{"I", "[I", "int", "java/lang/Integer"}
	Jlong    = &PrimitiveType{"J", "[J", "long", "java/lang/Long"}
	Jfloat   = &PrimitiveType{"F", "[F", "float", "java/lang/Float"}
	Jdouble  = &PrimitiveType{"D", "[D", "double", "java/lang/Double"}
)

// type jboolean bool
// type jbyte int8
// type jchar uint16
// type jshort int16
// type jint int32
// type jlong int64
// type jfloat float32
// type jdouble float64

type PrimitiveType struct {
	descriptor       string
	arrayClassName   string
	name             string
	wrapperClassName string
}
