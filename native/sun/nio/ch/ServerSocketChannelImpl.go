package ch

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
}

func _ssci(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/nio/ch/ServerSocketChannelImpl", name, desc, method)
}
