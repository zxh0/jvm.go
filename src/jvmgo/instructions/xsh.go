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

// // Boolean OR long
// type lor struct {NoOperandsInstruction}
// func (self *lor) execute(thread *rtda.Thread) {
//     stack := thread.CurrentFrame().OperandStack()
//     v1 := stack.PopLong()
//     v2 := stack.PopLong()
//     result := v1 | v2
//     stack.PushLong(result)
// }
