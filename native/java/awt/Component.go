package awt

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
}

func _comp(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Component", name, desc, method)
}
