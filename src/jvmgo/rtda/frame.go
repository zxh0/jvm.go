package rtda

import "jvmgo/rtda/class"

// stack frame
type Frame struct {
    nextPC          int
    localVars       *LocalVars
    operandStack    *OperandStack
    method          *class.Method
}

// getters & setters
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
func (self *Frame) Method() (*class.Method) {
    return self.method
}

func NewFrame(method *class.Method) (*Frame) {
    localVars := newLocalVars(method.MaxLocals())
    operandStack := newOperandStack(method.MaxStack())
    return &Frame{0, localVars, operandStack, method}
}
