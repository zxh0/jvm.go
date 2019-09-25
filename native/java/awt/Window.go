package awt

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func _window(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Window", name, desc, method)
}
