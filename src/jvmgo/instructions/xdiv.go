package instructions

import "jvmgo/rtda"

// Divide double
type ddiv struct {NoOperandsInstruction}
func (self *ddiv) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopDouble()
    v1 := stack.PopDouble()
    result := v1 / v2
    stack.PushDouble(result)
}

// Divide float
type fdiv struct {NoOperandsInstruction}
func (self *fdiv) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopFloat()
    v1 := stack.PopFloat()
    result := v1 / v2
    stack.PushFloat(result)
}

// Divide int
type idiv struct {NoOperandsInstruction}
func (self *idiv) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    result := v1 / v2
    stack.PushInt(result)
}

// Divide long
type ldiv struct {NoOperandsInstruction}
func (self *ldiv) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v2 := stack.PopLong()
    v1 := stack.PopLong()
    result := v1 / v2
    stack.PushLong(result)
}
