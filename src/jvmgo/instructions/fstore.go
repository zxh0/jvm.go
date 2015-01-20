package instructions

import "jvmgo/rtda"

// Store float into local variable 
type fstore struct {Index8Instruction}
func (self *fstore) Execute(thread *rtda.Thread) {
    _fstore(thread, uint(self.index))
}

type fstore_0 struct {NoOperandsInstruction}
func (self *fstore_0) Execute(thread *rtda.Thread) {
    _fstore(thread, 0)
}

type fstore_1 struct {NoOperandsInstruction}
func (self *fstore_1) Execute(thread *rtda.Thread) {
    _fstore(thread, 1)
}

type fstore_2 struct {NoOperandsInstruction}
func (self *fstore_2) Execute(thread *rtda.Thread) {
    _fstore(thread, 2)
}

type fstore_3 struct {NoOperandsInstruction}
func (self *fstore_3) Execute(thread *rtda.Thread) {
    _fstore(thread, 3)
}

func _fstore(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().OperandStack().PopFloat()
    thread.CurrentFrame().LocalVars().SetFloat(index, val)
}
