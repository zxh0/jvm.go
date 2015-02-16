package rtda

import (
	. "jvmgo/any"
	rtc "jvmgo/jvm/rtda/class"
)

func newShimFrame(thread *Thread, args []Any) *Frame {
	frame := &Frame{}
	frame.thread = thread
	frame.method = rtc.ReturnMethod()
	frame.operandStack = &OperandStack{
		size:  uint(len(args)),
		slots: args,
	}
	return frame
}

func newAthrowFrame(thread *Thread, ex *rtc.Obj, initArgs []Any) *Frame {
	// stackSlots := [ex, ex, initArgs]
	stackSlots := make([]Any, len(initArgs) + 2)
	stackSlots[0] = ex
	stackSlots[1] = ex
	copy(stackSlots[2:], initArgs)

	frame := &Frame{}
	frame.thread = thread
	frame.method = rtc.AthrowMethod()
	frame.operandStack = &OperandStack{
		size:  uint(len(stackSlots)),
		slots: stackSlots,
	}
	return frame
}
