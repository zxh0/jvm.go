package instructions

import "jvmgo/rtda"

// Store int into local variable 
type istore struct {Index8Instruction}
func (self *istore) execute(thread *rtda.Thread) {
    _istore(thread, uint(self.index))
}

type istore_0 struct {NoOperandsInstruction}
func (self *istore_0) execute(thread *rtda.Thread) {
    _istore(thread, 0)
}

type istore_1 struct {NoOperandsInstruction}
func (self *istore_1) execute(thread *rtda.Thread) {
    _istore(thread, 1)
}

type istore_2 struct {NoOperandsInstruction}
func (self *istore_2) execute(thread *rtda.Thread) {
    _istore(thread, 2)
}

type istore_3 struct {NoOperandsInstruction}
func (self *istore_3) execute(thread *rtda.Thread) {
    _istore(thread, 3)
}

func _istore(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    thread.CurrentFrame().LocalVars().SetInt(index, val)
}
