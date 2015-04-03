package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_frame(frame_initIDs, "initIDs", "()V")
}

func _frame(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Frame", name, desc, method)
}

func frame_initIDs(frame *rtda.Frame) {
	//TODO
}
