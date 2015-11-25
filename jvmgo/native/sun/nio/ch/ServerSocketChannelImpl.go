package ch

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
}

func _ssci(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("sun/nio/ch/ServerSocketChannelImpl", name, desc, method)
}
