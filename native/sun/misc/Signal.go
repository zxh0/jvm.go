package misc

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_signal(findSignal, "findSignal", "(Ljava/lang/String;)I")
	_signal(handle0, "handle0", "(IJ)J")
}

func _signal(method native.Method, name, desc string) {
	native.Register("sun/misc/Signal", name, desc, method)
}

// private static native int findSignal(String string);
// (Ljava/lang/String;)I
func findSignal(frame *rtda.Frame) {
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
