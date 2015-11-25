package io

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
	_i4a(i4a_init, "init", "()V")
}

func _i4a(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/net/Inet4Address", name, desc, method)
}

func i4a_init(frame *rtda.Frame) {

}
