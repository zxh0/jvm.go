package awt

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _cursor(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Cursor", name, desc, method)
}
