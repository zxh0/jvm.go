package instructions

import "jvmgo/jvm/rtda"

// Boolean OR int
type ior struct {NoOperandsInstruction}
func (self *ior) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopInt()
    v1 := stack.PopInt()
    result := v1 | v2
    stack.PushInt(result)
}

// Boolean OR long
type lor struct {NoOperandsInstruction}
func (self *lor) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    v2 := stack.PopLong()
    v1 := stack.PopLong()
    result := v1 | v2
    stack.PushLong(result)
}
