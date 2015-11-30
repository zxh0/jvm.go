package io

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
	_ia(ia_init, "init", "()V")
}

func _ia(method func(frame *rtda.Frame), name, desc string) {
	rtc.RegisterNativeMethod("java/net/InetAddress", name, desc, method)
}

func ia_init(frame *rtda.Frame) {

}
