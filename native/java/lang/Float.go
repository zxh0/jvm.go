package lang

import (
	"math"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/lang/Float").
		Register(floatToRawIntBits, "(F)I").
		Register(intBitsToFloat, "(I)F")
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.GetFloatVar(0)
	bits := math.Float32bits(value)

	frame.PushInt(int32(bits)) // todo
}

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *rtda.Frame) {
	bits := frame.GetIntVar(0)
	value := math.Float32frombits(uint32(bits))

	frame.PushFloat(value)
}
