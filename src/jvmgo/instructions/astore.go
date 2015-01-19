package instructions

import "jvmgo/rtda"

// Store reference into local variable 
type astore struct {Index8Instruction}
func (self *astore) execute(thread *rtda.Thread) {
    _astore(thread, uint(self.index))
}

type astore_0 struct {NoOperandsInstruction}
func (self *astore_0) execute(thread *rtda.Thread) {
    _astore(thread, 0)
}

type astore_1 struct {NoOperandsInstruction}
func (self *astore_1) execute(thread *rtda.Thread) {
    _astore(thread, 1)
}

type astore_2 struct {NoOperandsInstruction}
func (self *astore_2) execute(thread *rtda.Thread) {
    _astore(thread, 2)
}

type astore_3 struct {NoOperandsInstruction}
func (self *astore_3) execute(thread *rtda.Thread) {
    _astore(thread, 3)
}

func _astore(thread *rtda.Thread, index uint) {
    ref := thread.CurrentFrame().OperandStack().PopRef()
    thread.CurrentFrame().LocalVars().SetRef(index, ref)
}
