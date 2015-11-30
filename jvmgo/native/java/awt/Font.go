package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
}

func _font(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Font", name, desc, method)
}
