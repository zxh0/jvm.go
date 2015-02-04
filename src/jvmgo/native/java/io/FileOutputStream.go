package io

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _fos(fos_initIDs, "initIDs", "()V")
}

func _fos(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/io/FileOutputStream", name, desc, method)
}

// private static native void initIDs();
// ()V
func fos_initIDs(frame *rtda.Frame) {
    // todo
}
