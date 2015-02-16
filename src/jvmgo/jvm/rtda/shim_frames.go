package rtda

import (
    . "jvmgo/any"
    rtc "jvmgo/jvm/rtda/class"
)

func newReturnShimFrame(thread *Thread, args []Any) (*Frame) {
    frame := &Frame{}
    frame.thread = thread
    frame.method = rtc.ReturnShimMethod()
    frame.operandStack = &OperandStack{
        size: uint(len(args)),
        slots: args,
    }
    return frame
}
