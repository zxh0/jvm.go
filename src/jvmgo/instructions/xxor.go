package instructions

import "jvmgo/rtda"

// Boolean XOR int
type ixor struct {NoOperandsInstruction}
func (self *ixor) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 ^ v2
    stack.PushInt(result)
}

// Boolean XOR long
type lxor struct {NoOperandsInstruction}
func (self *lxor) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 ^ v2
    stack.PushLong(result)
}
