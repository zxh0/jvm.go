package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func _frame(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Frame", name, desc, method)
}
