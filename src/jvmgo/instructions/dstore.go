package instructions

import "jvmgo/rtda"

// Store double into local variable 
type dstore struct {Index8Instruction}
func (self *dstore) execute(thread *rtda.Thread) {
    _dstore(thread, uint(self.index))
}

type dstore_0 struct {NoOperandsInstruction}
func (self *dstore_0) execute(thread *rtda.Thread) {
    _dstore(thread, 0)
}

type dstore_1 struct {NoOperandsInstruction}
func (self *dstore_1) execute(thread *rtda.Thread) {
    _dstore(thread, 1)
}

type dstore_2 struct {NoOperandsInstruction}
func (self *dstore_2) execute(thread *rtda.Thread) {
    _dstore(thread, 2)
}

type dstore_3 struct {NoOperandsInstruction}
func (self *dstore_3) execute(thread *rtda.Thread) {
    _dstore(thread, 3)
}

func _dstore(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().OperandStack().PopDouble()
    thread.CurrentFrame().LocalVars().SetDouble(index, val)
}
