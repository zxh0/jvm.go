package misc

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_signal(findSignal0, "findSignal0", "(Ljava/lang/String;)I")
	_signal(handle0, "handle0", "(IJ)J")
}

func _signal(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("jdk/internal/misc/Signal", name, desc, method)
}

// private static native int findSignal0(String string);
// (Ljava/lang/String;)I
func findSignal0(frame *rtda.Frame) {
	frame.GetRefVar(0) // name

	frame.PushInt(0) // todo
}

// private static native long handle0(int i, long l);
// (IJ)J
func handle0(frame *rtda.Frame) {
	// todo
	frame.GetIntVar(0)
	frame.GetLongVar(1)

	frame.PushLong(0)
}

// private static native void raise0(int i);
