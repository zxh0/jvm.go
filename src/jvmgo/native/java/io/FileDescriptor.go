package io

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_fd(fd_initIDs, "initIDs", "()V")
}

func _fd(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/FileDescriptor", name, desc, method)
}

// private static native void initIDs();
// ()V
func fd_initIDs(frame *rtda.Frame) {
	// todo
}
