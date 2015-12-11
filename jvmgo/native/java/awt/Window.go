package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func _window(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Window", name, desc, method)
}
