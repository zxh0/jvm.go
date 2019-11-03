package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_ia(ia_init, "init", "()V")
}

func _ia(method native.Method, name, desc string) {
	native.Register("java/net/InetAddress", name, desc, method)
}

func ia_init(frame *rtda.Frame) {

}
