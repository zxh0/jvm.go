package ref

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_ref(getAndClearReferencePendingList, "getAndClearReferencePendingList", "()Ljava/lang/ref/Reference;")
	_ref(hasReferencePendingList, "hasReferencePendingList", "()B")
	_ref(waitForReferencePendingList, "waitForReferencePendingList", "()V")
}

func _ref(method native.Method, name, desc string) {
	native.Register("java/lang/ref/Reference", name, desc, method)
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
