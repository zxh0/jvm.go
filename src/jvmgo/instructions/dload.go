package instructions

import "jvmgo/rtda"

// Load double from local variable 
type dload struct {Index8Instruction}
func (self *dload) Execute(thread *rtda.Thread) {
    _dload(thread, uint(self.index))
}

type dload_0 struct {NoOperandsInstruction}
func (self *dload_0) Execute(thread *rtda.Thread) {
    _dload(thread, 0)
}

type dload_1 struct {NoOperandsInstruction}
func (self *dload_1) Execute(thread *rtda.Thread) {
    _dload(thread, 1)
}

type dload_2 struct {NoOperandsInstruction}
func (self *dload_2) Execute(thread *rtda.Thread) {
    _dload(thread, 2)
}

type dload_3 struct {NoOperandsInstruction}
func (self *dload_3) Execute(thread *rtda.Thread) {
    _dload(thread, 3)
}

func _dload(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().LocalVars().GetDouble(index)
    thread.CurrentFrame().OperandStack().PushDouble(val)
}
