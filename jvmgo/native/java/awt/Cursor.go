package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _cursor(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Cursor", name, desc, method)
}
