package awt

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func _frame(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Frame", name, desc, method)
}
