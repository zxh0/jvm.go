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

// Duplicate the top operand stack value and insert two values down
type dup_x1 struct {}
func (self *dup_x1) fetchOperands(bcr *BytecodeReader) {}
func (self *dup_x1) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.Pop()
    val2 := stack.Pop()
    stack.Push(val1)
    stack.Push(val2)
    stack.Push(val1)
}
