package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
}

func _font(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Font", name, desc, method)
}
