package io

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_osc(initNative, "initNative", "()V")
}

func _osc(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/ObjectStreamClass", name, desc, method)
}

// private static native void initNative();
// ()V
func initNative(frame *rtda.Frame) {
	// todo
}
