package misc

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("jdk/internal/misc/VM").
		Register(initialize, "()V").
		Register(initializeFromArchive, "(Ljava/lang/Class;)V")
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
