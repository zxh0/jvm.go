package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
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
