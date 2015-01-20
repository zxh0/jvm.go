package instructions

import "jvmgo/rtda"

// Multiply double
type dmul struct {NoOperandsInstruction}
func (self *dmul) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 * v2
    stack.PushDouble(result)
}

// Multiply float
type fmul struct {NoOperandsInstruction}
func (self *fmul) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := v1 * v2
    stack.PushFloat(result)
}

// Multiply int
type imul struct {NoOperandsInstruction}
func (self *imul) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 * v2
    stack.PushInt(result)
}

// Multiply long
type lmul struct {NoOperandsInstruction}
func (self *lmul) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 * v2
    stack.PushLong(result)
}
