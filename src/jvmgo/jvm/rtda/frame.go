package rtda

import rtc "jvmgo/rtda/class"

// stack frame
type Frame struct {
    thread          *Thread
    nextPC          int
    localVars       *LocalVars
    operandStack    *OperandStack
    method          *rtc.Method
    onPopAction     func()
}

func newFrame(thread *Thread, method *rtc.Method) (*Frame) {
    localVars := newLocalVars(method.MaxLocals())
    operandStack := newOperandStack(method.MaxStack())
    return &Frame{thread, 0, localVars, operandStack, method, nil}
}

// getters & setters
func (self *Frame) Thread() (*Thread) {
    return self.thread
}
func (self *Frame) NextPC() (int) {
    return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
    self.nextPC = nextPC
}
func (self *Frame) LocalVars() (*LocalVars) {
    return self.localVars
}
func (self *Frame) OperandStack() (*OperandStack) {
    return self.operandStack
}
func (self *Frame) Method() (*rtc.Method) {
    return self.method
}
func (self *Frame) SetOnPopAction(f func()) {
    self.onPopAction = f
}
