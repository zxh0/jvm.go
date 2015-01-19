package instructions

import "jvmgo/rtda"

// Load int from local variable 
type iload struct {Index8Instruction}
func (self *iload) execute(thread *rtda.Thread) {
    _iload(thread, uint(self.index))
}

type iload_0 struct {NoOperandsInstruction}
func (self *iload_0) execute(thread *rtda.Thread) {
    _iload(thread, 0)
}

type iload_1 struct {NoOperandsInstruction}
func (self *iload_1) execute(thread *rtda.Thread) {
    _iload(thread, 1)
}

type iload_2 struct {NoOperandsInstruction}
func (self *iload_2) execute(thread *rtda.Thread) {
    _iload(thread, 2)
}

type iload_3 struct {NoOperandsInstruction}
func (self *iload_3) execute(thread *rtda.Thread) {
    _iload(thread, 3)
}

func _iload(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().LocalVars().GetInt(index)
    thread.CurrentFrame().OperandStack().PushInt(val)
}
