package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Get length of array
type arraylength struct {NoOperandsInstruction}
func (self *arraylength) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    arrRef := stack.PopRef()
    arrLen := class.ArrayLength(arrRef)
    stack.PushInt(arrLen)
}
