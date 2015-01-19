package instructions

import "jvmgo/rtda"

// Load long from local variable 
type lload struct {Index8Instruction}
func (self *lload) execute(thread *rtda.Thread) {
    _lload(thread, uint(self.index))
}

type lload_0 struct {NoOperandsInstruction}
func (self *lload_0) execute(thread *rtda.Thread) {
    _lload(thread, 0)
}

type lload_1 struct {NoOperandsInstruction}
func (self *lload_1) execute(thread *rtda.Thread) {
    _lload(thread, 1)
}

type lload_2 struct {NoOperandsInstruction}
func (self *lload_2) execute(thread *rtda.Thread) {
    _lload(thread, 2)
}

type lload_3 struct {NoOperandsInstruction}
func (self *lload_3) execute(thread *rtda.Thread) {
    _lload(thread, 3)
}

func _lload(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().LocalVars().GetLong(index)
    thread.CurrentFrame().OperandStack().PushLong(val)
}
