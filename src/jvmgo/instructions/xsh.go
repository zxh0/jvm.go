package instructions

import "jvmgo/rtda"

// Shift left int
type ishl struct {NoOperandsInstruction}
func (self *ishl) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x1f
    result := v1 << s
    stack.PushInt(result)
}

// Arithmetic shift right int
type ishr struct {NoOperandsInstruction}
func (self *ishr) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x1f
    result := v1 >> s // todo
    stack.PushInt(result)
}

// Logical shift right int
type iushr struct {NoOperandsInstruction}
func (self *iushr) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    s := uint32(v2) & 0x1f
    result := v1 >> s
    stack.PushInt(result)
}
