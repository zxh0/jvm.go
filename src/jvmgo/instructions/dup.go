package instructions

import "jvmgo/rtda"

// Duplicate the top operand stack value
type dup struct {}
func (self *dup) fetchOperands(bcr *BytecodeReader) {}
func (self *dup) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.Pop()
    stack.Push(val)
    stack.Push(val)
}
