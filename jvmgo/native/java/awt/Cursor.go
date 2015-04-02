package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_cursor(cursor_initIDs, "initIDs", "()V")
}

func _cursor(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Cursor", name, desc, method)
}

func cursor_initIDs(frame *rtda.Frame) {
	//TODO
}
