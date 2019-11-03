package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_fd(setErrorMode, "setErrorMode", "(J)J")
}

func _fd(method native.Method, name, desc string) {
	native.Register("sun/io/Win32ErrorMode", name, desc, method)
}

func setErrorMode(frame *rtda.Frame) {
	frame.PushLong(0)
}
