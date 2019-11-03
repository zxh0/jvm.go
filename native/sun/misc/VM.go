package misc

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_vm(initialize, "initialize", "()V")
}

func _vm(method native.Method, name, desc string) {
	native.Register("sun/misc/VM", name, desc, method)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
	// todo
}
