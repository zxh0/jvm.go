package ch

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _ssci(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/nio/ch/ServerSocketChannelImpl", name, desc, method)
}
