package rtda

import (
	. "jvmgo/any"
	rtc "jvmgo/jvm/rtda/class"
)

func newShimFrame(thread *Thread, args []Any) *Frame {
	frame := &Frame{}
	frame.thread = thread
	frame.method = rtc.ShimMethod()
	frame.operandStack = &OperandStack{
		size:  uint(len(args)),
		slots: args,
	}
	return frame
}
