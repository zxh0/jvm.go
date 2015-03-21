package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_id(id_init, "init", "()V")
}

func _id(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/InetAddress", name, desc, method)
}

func id_init(frame *rtda.Frame) {

}
