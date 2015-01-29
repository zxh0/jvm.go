package instructions

import "jvmgo/rtda"

// Subtract double
type dsub struct {NoOperandsInstruction}
func (self *dsub) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopDouble()
    v1 := stack.PopDouble()
    result := v1 - v2
    stack.PushDouble(result)
}

// Subtract float
type fsub struct {NoOperandsInstruction}
func (self *fsub) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopFloat()
    v1 := stack.PopFloat()
    result := v1 - v2
    stack.PushFloat(result)
}

// Subtract int
type isub struct {NoOperandsInstruction}
func (self *isub) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    result := v1 - v2
    stack.PushInt(result)
}

// Subtract long
type lsub struct {NoOperandsInstruction}
func (self *lsub) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopLong()
    v1 := stack.PopLong()
    result := v1 - v2
    stack.PushLong(result)
}
