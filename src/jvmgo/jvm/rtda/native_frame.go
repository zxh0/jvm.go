package rtda

import rtc "jvmgo/jvm/rtda/class"

func newNativeFrame(thread *Thread, method *rtc.Method) (*Frame) {
    frame := &Frame{}
    frame.thread = thread
    frame.method = method
    frame.localVars = newLocalVars(method.ActualArgCount())
    frame.operandStack = newOperandStack(method.MaxStack())
    return frame
}
