package instructions

import "jvmgo/rtda"

// Shift left int
type ishl struct {NoOperandsInstruction}
func (self *ishl) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := uint32(stack.PopInt()) & 0x1f
    result := v1 << v2
    stack.PushInt(result)
}

// Arithmetic shift right int
type ishr struct {NoOperandsInstruction}
func (self *ishr) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := uint32(stack.PopInt()) & 0x1f
    result := v1 >> v2 // todo
    stack.PushInt(result)
}
