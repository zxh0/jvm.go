package instructions

import "jvmgo/rtda"

// Boolean OR int
type ior struct {NoOperandsInstruction}
func (self *ior) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 | v2
    stack.PushInt(result)
}

// Boolean OR long
type lor struct {NoOperandsInstruction}
func (self *lor) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 | v2
    stack.PushLong(result)
}
