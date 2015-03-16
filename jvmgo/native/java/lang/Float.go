package lang

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"math"
)

func init() {
	_float(floatToRawIntBits, "floatToRawIntBits", "(F)I")
}

func _float(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/Float", name, desc, method)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *rtda.Frame) {
	vars := frame.LocalVars()
	value := vars.GetFloat(0)
	bits := math.Float32bits(value)

	stack := frame.OperandStack()
	stack.PushInt(int32(bits)) // todo
}
