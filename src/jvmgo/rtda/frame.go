package rtda

// stack frame
type Frame struct {
    localVars       *LocalVars
    operandStack    *OperandStack
}

func newFrame(localVarsSize uint16, operandStackSize uint16) (*Frame) {
    localVars := newLocalVars(localVarsSize)
    operandStack := newOperandStack(operandStackSize)
    return &Frame{localVars, operandStack}
}
