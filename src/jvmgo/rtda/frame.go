package rtda

import "jvmgo/rtda/class"

// stack frame
type Frame struct {
    nextPC          int
    localVars       *LocalVars
    operandStack    *OperandStack
    method          *class.Method
}

// getters
func (self *Frame) LocalVars() (*LocalVars) {
    return self.localVars
}
func (self *Frame) OperandStack() (*OperandStack) {
    return self.operandStack
}
func (self *Frame) Method() (*class.Method) {
    return self.method
}

func newFrame(localVarsSize, operandStackSize uint16) (*Frame) {
    localVars := newLocalVars(localVarsSize)
    operandStack := newOperandStack(operandStackSize)
    return &Frame{0, localVars, operandStack, nil} // todo
}
