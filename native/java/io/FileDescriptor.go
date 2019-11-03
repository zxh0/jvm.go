package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_fd(set, "set", "(I)J")
}

func _fd(method native.Method, name, desc string) {
	native.Register("java/io/FileDescriptor", name, desc, method)
}

func set(frame *rtda.Frame) {
	frame.PushLong(0)
}
