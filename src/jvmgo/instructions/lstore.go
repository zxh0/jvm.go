package instructions

import "jvmgo/rtda"

// Store long into local variable 
type lstore struct {Index8Instruction}
func (self *lstore) execute(thread *rtda.Thread) {
    _lstore(thread, uint(self.index))
}

type lstore_0 struct {NoOperandsInstruction}
func (self *lstore_0) execute(thread *rtda.Thread) {
    _lstore(thread, 0)
}

type lstore_1 struct {NoOperandsInstruction}
func (self *lstore_1) execute(thread *rtda.Thread) {
    _lstore(thread, 1)
}

type lstore_2 struct {NoOperandsInstruction}
func (self *lstore_2) execute(thread *rtda.Thread) {
    _lstore(thread, 2)
}

type lstore_3 struct {NoOperandsInstruction}
func (self *lstore_3) execute(thread *rtda.Thread) {
    _lstore(thread, 3)
}

func _lstore(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().OperandStack().PopLong()
    thread.CurrentFrame().LocalVars().SetLong(index, val)
}
