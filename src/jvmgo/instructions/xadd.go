package instructions

import "jvmgo/rtda"

// Add double
type dadd struct {NoOperandsInstruction}
func (self *dadd) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 + v2
    stack.PushDouble(result)
}

// Add float
type fadd struct {NoOperandsInstruction}
func (self *fadd) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopFloat()
    v1 := stack.PopFloat()
    result := v1 + v2
    stack.PushFloat(result)
}

// Add int
type iadd struct {NoOperandsInstruction}
func (self *iadd) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    result := v1 + v2
    stack.PushInt(result)
}

// Add long
type ladd struct {NoOperandsInstruction}
func (self *ladd) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopLong()
    v1 := stack.PopLong()
    result := v1 + v2
    stack.PushLong(result)
}
