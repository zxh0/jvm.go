package io

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _ufs(ufs_initIDs, "initIDs", "()V")
}

func _ufs(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/io/UnixFileSystem", name, desc, method)
}

// private static native void initIDs();
// ()V
func ufs_initIDs(frame *rtda.Frame) {
    // todo
}
