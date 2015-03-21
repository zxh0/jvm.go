package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_psi(psi_initProto, "initProto", "()V")
}

func _psi(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/PlainSocketImpl", name, desc, method)
}

func psi_initProto(frame *rtda.Frame) {
	//TODO
}
