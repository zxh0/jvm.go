package io

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_fd(set, "set", "(I)J")
}

func _fd(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/FileDescriptor", name, desc, method)
}

func set(frame *rtda.Frame) {
	frame.PushLong(0)
}
