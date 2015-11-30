package ch

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
}

func _ssci(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("sun/nio/ch/ServerSocketChannelImpl", name, desc, method)
}
