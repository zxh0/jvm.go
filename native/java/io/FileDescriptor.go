package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_fd(set, "set", "(I)J")
	_fd(getHandle, "getHandle", "(I)J")
	_fd(getAppend, "getAppend", "(I)Z")
}

func _fd(method native.Method, name, desc string) {
	native.Register("java/io/FileDescriptor", name, desc, method)
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
