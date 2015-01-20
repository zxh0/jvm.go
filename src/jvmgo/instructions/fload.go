package instructions

import "jvmgo/rtda"

// Load float from local variable 
type fload struct {Index8Instruction}
func (self *fload) Execute(thread *rtda.Thread) {
    _fload(thread, uint(self.index))
}

type fload_0 struct {NoOperandsInstruction}
func (self *fload_0) Execute(thread *rtda.Thread) {
    _fload(thread, 0)
}

type fload_1 struct {NoOperandsInstruction}
func (self *fload_1) Execute(thread *rtda.Thread) {
    _fload(thread, 1)
}

type fload_2 struct {NoOperandsInstruction}
func (self *fload_2) Execute(thread *rtda.Thread) {
    _fload(thread, 2)
}

type fload_3 struct {NoOperandsInstruction}
func (self *fload_3) Execute(thread *rtda.Thread) {
    _fload(thread, 3)
}

func _fload(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().LocalVars().GetFloat(index)
    thread.CurrentFrame().OperandStack().PushFloat(val)
}
