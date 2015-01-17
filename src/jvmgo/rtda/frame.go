package rtda

// stack frame
type Frame struct {
    localVars       *LocalVars
    operandStack    *OperandStack
}

func newFrame(localVarsLen uint16, operandStackLen uint16) (*Frame) {
    localVars := newLocalVars(localVarsLen)
    operandStack := newOperandStack(operandStackLen)
    return &Frame{localVars, operandStack}
}
