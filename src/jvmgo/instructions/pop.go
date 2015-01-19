package instructions

import "jvmgo/rtda"

// Pop the top operand stack value
type pop struct {NoOperandsInstruction}
func (self *pop) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().Pop()
}
