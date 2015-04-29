package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_fd(fd_set, "set", "(I)J")
}

func _fd(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/FileDescriptor", name, desc, method)
}

func fd_set(frame *rtda.Frame) {
	//TODO
	vars := frame.LocalVars()
	fd := vars.GetInt(0)
	stack := frame.OperandStack()
	stack.PushLong(int64(fd))
}
