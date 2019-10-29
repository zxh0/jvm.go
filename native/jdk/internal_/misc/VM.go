package misc

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_vm(initialize, "initialize", "()V")
	_vm(initializeFromArchive, "initializeFromArchive", "(Ljava/lang/Class;)V")
}

func _vm(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("jdk/internal/misc/VM", name, desc, method)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {
	// TODO
}

// public static native void initializeFromArchive(Class<?> c);
// (Ljava/lang/Class;)V
func initializeFromArchive(frame *rtda.Frame) {
	// TODO
}
