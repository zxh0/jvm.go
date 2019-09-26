package io

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_fd(setErrorMode, "setErrorMode", "(J)J")
}

func _fd(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/io/Win32ErrorMode", name, desc, method)
}

func setErrorMode(frame *rtda.Frame) {
	frame.PushLong(0)
}
