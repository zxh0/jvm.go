package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func _frame(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Frame", name, desc, method)
}
