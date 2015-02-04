package io

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _fis(initIDs, "initIDs", "()V")
}

func _fis(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/io/FileInputStream", name, desc, method)
}

// private static native void initIDs();
// ()V
func initIDs(frame *rtda.Frame) {
    // todo
}
