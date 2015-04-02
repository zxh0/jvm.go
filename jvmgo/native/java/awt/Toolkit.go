package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_tk(tk_initIDs, "initIDs", "()V")
}

func _tk(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Toolkit", name, desc, method)
}

func tk_initIDs(frame *rtda.Frame) {
	//TODO
}
