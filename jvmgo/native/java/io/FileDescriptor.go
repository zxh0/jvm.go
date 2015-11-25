package io

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
}

func _fd(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/io/FileDescriptor", name, desc, method)
}
