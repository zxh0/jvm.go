package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_i4a(i4a_init, "init", "()V")
}

func _i4a(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/Inet4Address", name, desc, method)
}

func i4a_init(frame *rtda.Frame) {

}
