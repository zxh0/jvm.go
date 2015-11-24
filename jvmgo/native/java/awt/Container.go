package awt

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _container(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Container", name, desc, method)
}
