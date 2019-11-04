package ref

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/lang/ref/Reference").
		Register(getAndClearReferencePendingList, "()Ljava/lang/ref/Reference;").
		Register(hasReferencePendingList, "()B").
		Register(waitForReferencePendingList, "()V")
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
