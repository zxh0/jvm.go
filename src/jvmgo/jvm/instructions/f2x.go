package instructions

import "jvmgo/jvm/rtda"

// Convert float to double
type f2d struct {NoOperandsInstruction}
func (self *f2d) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    f := stack.PopFloat()
    d := float64(f)
    stack.PushDouble(d)
}

// Convert float to int
type f2i struct {NoOperandsInstruction}
func (self *f2i) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    f := stack.PopFloat()
    i := int32(f)
    stack.PushInt(i)
}

// Convert float to long
type f2l struct {NoOperandsInstruction}
func (self *f2l) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    f := stack.PopFloat()
    l := int64(f)
    stack.PushLong(l)
}
