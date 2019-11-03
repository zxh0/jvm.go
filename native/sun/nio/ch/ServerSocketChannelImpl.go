package ch

import (
	"github.com/zxh0/jvm.go/native"
)

func init() {
}

func _ssci(method native.Method, name, desc string) {
	native.Register("sun/nio/ch/ServerSocketChannelImpl", name, desc, method)
}
