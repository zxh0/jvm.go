package io

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_fis(fis_initIDs, "initIDs", "()V")
}

func _fis(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/FileInputStream", name, desc, method)
}

// private static native void initIDs();
// ()V
func fis_initIDs(frame *rtda.Frame) {
	// todo
}
