package misc

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_vm(initialize, "initialize", "()V")
}

func _vm(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/misc/VM", name, desc, method)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
	// todo
}
