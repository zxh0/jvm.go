package lang

import (
    "math"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _float(floatToRawIntBits, "floatToRawIntBits", "(F)I")
}

func _float(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Float", name, desc, method)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *rtda.Frame, x int) {
    vars := frame.LocalVars()
    value := vars.GetFloat(0)
    bits := math.Float32bits(value)

    stack := frame.OperandStack()
    stack.PushInt(int32(bits)) // todo
}
