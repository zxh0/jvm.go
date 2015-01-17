package rtda

// stack frame
type Frame struct {
    localVars       *LocalVars
    operandStack    *OperandStack
}

func (self *Frame) executeOneBytecode() {
    // todo
}

func newFrame(localVarsSize, operandStackSize uint16) (*Frame) {
    localVars := newLocalVars(localVarsSize)
    operandStack := newOperandStack(operandStackSize)
    return &Frame{localVars, operandStack}
}
