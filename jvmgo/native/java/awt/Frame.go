package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func _frame(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Frame", name, desc, method)
}
