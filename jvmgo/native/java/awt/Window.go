package awt

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func _window(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Window", name, desc, method)
}
