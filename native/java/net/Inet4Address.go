package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_i4a(i4a_init, "init", "()V")
}

func _i4a(method native.Method, name, desc string) {
	native.Register("java/net/Inet4Address", name, desc, method)
}

func i4a_init(frame *rtda.Frame) {

}
