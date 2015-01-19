package instructions

import "jvmgo/rtda"

// Convert long to double
type l2d struct {NoOperandsInstruction}
func (self *l2d) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    l := stack.PopLong()
    d := float64(l)
    stack.PushDouble(d)
}

// Convert long to float
type l2f struct {NoOperandsInstruction}
func (self *l2f) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    l := stack.PopLong()
    f := float32(l)
    stack.PushFloat(f)
}

// Convert long to int
type l2i struct {NoOperandsInstruction}
func (self *l2i) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    l := stack.PopLong()
    i := int32(l)
    stack.PushInt(i)
}
