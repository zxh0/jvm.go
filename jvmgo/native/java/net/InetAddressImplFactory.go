package io

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_iaif(iaif_isIPv6Supported, "isIPv6Supported", "()Z")
}

func _iaif(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/net/InetAddressImplFactory", name, desc, method)
}

//static native boolean isIPv6Supported();
// ()Z
func iaif_isIPv6Supported(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(true)
}
