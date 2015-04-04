package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _tk(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Toolkit", name, desc, method)
}
