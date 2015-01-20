package instructions

import "jvmgo/rtda"

// Convert double to float
type d2f struct {NoOperandsInstruction}
func (self *d2f) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    d := stack.PopDouble()
    f := float32(d)
    stack.PushFloat(f)
}

// Convert double to int
type d2i struct {NoOperandsInstruction}
func (self *d2i) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    d := stack.PopDouble()
    i := int32(d)
    stack.PushInt(i)
}

// Convert double to long
type d2l struct {NoOperandsInstruction}
func (self *d2l) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    d := stack.PopDouble()
    l := int64(d)
    stack.PushLong(l)
}
