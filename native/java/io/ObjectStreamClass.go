package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/io/ObjectStreamClass").
		Register(initNative, "()V")
}

// private static native void initNative();
// ()V
func initNative(frame *rtda.Frame) {
	// todo
}
