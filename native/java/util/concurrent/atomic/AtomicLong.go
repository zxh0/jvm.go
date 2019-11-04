package atomic

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/util/concurrent/atomic/AtomicLong").
		Register(VMSupportsCS8, "()Z")
}

// private static native boolean VMSupportsCS8();
// ()Z
func VMSupportsCS8(frame *rtda.Frame) {
	frame.PushBoolean(false) // todo sync/atomic
}
