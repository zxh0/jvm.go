package ref

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_ref(getAndClearReferencePendingList, "getAndClearReferencePendingList", "()Ljava/lang/ref/Reference;")
	_ref(hasReferencePendingList, "hasReferencePendingList", "()B")
	_ref(waitForReferencePendingList, "waitForReferencePendingList", "()V")
}

func _ref(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/ref/Reference", name, desc, method)
}

// private static native Reference<Object> getAndClearReferencePendingList();
func getAndClearReferencePendingList(frame *rtda.Frame) {
	// TODO
	frame.PushNull()
}

// private static native boolean hasReferencePendingList();
func hasReferencePendingList(frame *rtda.Frame) {
	panic("TODO")
}

// private static native void waitForReferencePendingList();
func waitForReferencePendingList(frame *rtda.Frame) {
	// TODO
}
