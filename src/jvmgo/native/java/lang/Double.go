package lang

import (
    "math"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _double(doubleToRawLongBits, "doubleToRawLongBits", "(D)J")
}

func _double(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Double", name, desc, method)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtda.Frame, x int) {
    vars := frame.LocalVars()
    value := vars.GetDouble(0)
    bits := math.Float64bits(value)

    stack := frame.OperandStack()
    stack.PushLong(int64(bits)) // todo
}
