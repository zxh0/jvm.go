package misc

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _vm(initialize, "initialize", "()V")
}

func _vm(method Any, name, desc string) {
    rtc.RegisterNativeMethod("sun/misc/VM", name, desc, method)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
    // todo
}
