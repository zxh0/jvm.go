package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Get length of array
type arraylength struct {NoOperandsInstruction}
func (self *arraylength) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    arrRef := stack.PopRef()
    arrLen := class.ArrayLength(arrRef)
    stack.PushInt(arrLen)
}
