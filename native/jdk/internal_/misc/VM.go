package misc

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_vm(initialize, "initialize", "()V")
	_vm(initializeFromArchive, "initializeFromArchive", "(Ljava/lang/Class;)V")
}

func _vm(method native.Method, name, desc string) {
	native.Register("jdk/internal/misc/VM", name, desc, method)
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
