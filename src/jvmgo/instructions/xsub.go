package instructions

import "jvmgo/rtda"

// Subtract double
type dsub struct {NoOperandsInstruction}
func (self *dsub) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 - v2
    stack.PushDouble(result)
}

// Subtract float
type fsub struct {NoOperandsInstruction}
func (self *fsub) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := v1 - v2
    stack.PushFloat(result)
}

// Subtract int
type isub struct {NoOperandsInstruction}
func (self *isub) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 - v2
    stack.PushInt(result)
}

// Subtract long
type lsub struct {NoOperandsInstruction}
func (self *lsub) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 - v2
    stack.PushLong(result)
}
