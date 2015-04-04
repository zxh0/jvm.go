package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
}

func _fd(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/FileDescriptor", name, desc, method)
}
