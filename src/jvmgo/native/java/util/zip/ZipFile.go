package zip

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_zf(initIDs, "initIDs", "()V")
}

func _zf(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/util/zip/ZipFile", name, desc, method)
}

// private static native void initIDs();
// ()V
func initIDs(frame *rtda.Frame) {
	// todo
}
