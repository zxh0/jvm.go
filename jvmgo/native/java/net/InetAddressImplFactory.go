package io

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_iaif(iaif_isIPv6Supported, "isIPv6Supported", "()Z")
}

func _iaif(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/net/InetAddressImplFactory", name, desc, method)
}

//static native boolean isIPv6Supported();
// ()Z
func iaif_isIPv6Supported(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(true)
}
