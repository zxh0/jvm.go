package instructions

import "jvmgo/rtda"

// Pop the top operand stack value
type pop struct {NoOperandsInstruction}
func (self *pop) Execute(frame *rtda.Frame) {
    frame.OperandStack().Pop()
}

// Pop the top one or two operand stack values
type pop2 struct {NoOperandsInstruction}
func (self *pop2) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    val1 := stack.Pop()
    if !isLongOrDouble(val1) {
        stack.Pop()
    }
}
