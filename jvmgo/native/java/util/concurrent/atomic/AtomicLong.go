package atomic

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_al(VMSupportsCS8, "VMSupportsCS8", "()Z")
}

func _al(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/util/concurrent/atomic/AtomicLong", name, desc, method)
}

// private static native boolean VMSupportsCS8();
// ()Z
func VMSupportsCS8(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushBoolean(false) // todo sync/atomic
}
