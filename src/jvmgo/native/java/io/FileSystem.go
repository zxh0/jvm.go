package io

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _fs(getFileSystem, "getFileSystem", "()Ljava/io/FileSystem;")
}

func _fs(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/io/FileSystem", name, desc, method)
}

// public static native FileSystem getFileSystem()
// ()Ljava/io/FileSystem;
func getFileSystem(frame *rtda.Frame) {
    // todo
    panic("getFileSystem!!")
}
