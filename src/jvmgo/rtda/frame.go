package rtda

import "jvmgo/rtda/class"

// stack frame
type Frame struct {
    localVars       *LocalVars
    operandStack    *OperandStack
    method          *class.Class
}

func (self *Frame) LocalVars() (*LocalVars) {
    return self.localVars
}

func (self *Frame) OperandStack() (*OperandStack) {
    return self.operandStack
}

func (self *Frame) executeOneInstruction() {
    // todo
}

func newFrame(localVarsSize, operandStackSize uint16) (*Frame) {
    localVars := newLocalVars(localVarsSize)
    operandStack := newOperandStack(operandStackSize)
    return &Frame{localVars, operandStack, nil} // todo
}
