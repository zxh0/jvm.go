package misc

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("jdk/internal/misc/Signal").
		Register(findSignal0, "(Ljava/lang/String;)I").
		Register(handle0, "(IJ)J")
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
