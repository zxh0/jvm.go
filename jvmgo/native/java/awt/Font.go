package awt

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _font(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Font", name, desc, method)
}
