package lang

import (
    "math"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
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
    stack := frame.OperandStack()
    value := stack.PopFloat()
    bits := math.Float32bits(value)
    stack.PushInt(int32(bits)) // todo
}
