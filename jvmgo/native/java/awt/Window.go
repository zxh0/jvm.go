package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_window(window_initIDs, "initIDs", "()V")
}

func _window(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Window", name, desc, method)
}

func window_initIDs(frame *rtda.Frame) {
	//TODO
}
