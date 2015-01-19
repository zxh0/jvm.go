package instructions

import "jvmgo/rtda"

// Negate double
type dneg struct {NoOperandsInstruction}
func (self *dneg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopDouble()
    stack.PushDouble(-val)
}

// Negate float
type fneg struct {NoOperandsInstruction}
func (self *fneg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopFloat()
    stack.PushFloat(-val)
}

// Negate int
type ineg struct {NoOperandsInstruction}
func (self *ineg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    stack.PushInt(-val)
}

// Negate long
type lneg struct {NoOperandsInstruction}
func (self *lneg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopLong()
    stack.PushLong(-val)
}
