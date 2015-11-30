package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
}

func _comp(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Component", name, desc, method)
}
