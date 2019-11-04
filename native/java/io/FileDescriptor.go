package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/io/FileDescriptor").
		Register(set, "(I)J").
		Register(getHandle, "(I)J").
		Register(getAppend, "(I)Z")
}

func set(frame *rtda.Frame) {
	frame.PushLong(0)
}

// private static native long getHandle(int d);
func getHandle(frame *rtda.Frame) {
	// TODO
	frame.PushLong(0)
}

// private static native boolean getAppend(int fd);
func getAppend(frame *rtda.Frame) {
	// TODO
	frame.PushBoolean(false)
}
