package io

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_i4a(i4a_init, "init", "()V")
}

func _i4a(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/net/Inet4Address", name, desc, method)
}

func i4a_init(frame *rtda.Frame) {

}
