package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_comp(comp_initIDs, "initIDs", "()V")
}

func _comp(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Component", name, desc, method)
}

func comp_initIDs(frame *rtda.Frame) {
	//TODO
}
