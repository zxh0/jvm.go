package zip

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
	_inflater(inflater_initIDs, "initIDs", "()V")
}

func _inflater(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("java/util/zip/Inflater", name, desc, method)
}

// private static native void initIDs();
// ()V
func inflater_initIDs(frame *rtda.Frame) {
	// todo
}
