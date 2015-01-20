package instructions

import "jvmgo/rtda"

// Load reference from local variable 
type aload struct {Index8Instruction}
func (self *aload) Execute(thread *rtda.Thread) {
    _aload(thread, uint(self.index))
}

type aload_0 struct {NoOperandsInstruction}
func (self *aload_0) Execute(thread *rtda.Thread) {
    _aload(thread, 0)
}

type aload_1 struct {NoOperandsInstruction}
func (self *aload_1) Execute(thread *rtda.Thread) {
    _aload(thread, 1)
}

type aload_2 struct {NoOperandsInstruction}
func (self *aload_2) Execute(thread *rtda.Thread) {
    _aload(thread, 2)
}

type aload_3 struct {NoOperandsInstruction}
func (self *aload_3) Execute(thread *rtda.Thread) {
    _aload(thread, 3)
}

func _aload(thread *rtda.Thread, index uint) {
    ref := thread.CurrentFrame().LocalVars().GetRef(index)
    thread.CurrentFrame().OperandStack().PushRef(ref)
}
