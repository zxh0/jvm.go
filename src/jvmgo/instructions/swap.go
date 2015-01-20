package instructions

import "jvmgo/rtda"

// Swap the top two operand stack values
type swap struct {NoOperandsInstruction}
func (self *swap) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.Pop()
    val2 := stack.Pop()
    stack.Push(val1)
    stack.Push(val2)
}
