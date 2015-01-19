package instructions

import "jvmgo/rtda"

// Load reference from local variable 
type aload struct {
    index uint8
}
func (self *aload) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint8()
}
func (self *aload) execute(thread *rtda.Thread) {
    _aload(thread, uint(self.index))
}

type aload_0 struct {NoOperandsInstruction}
func (self *aload_0) execute(thread *rtda.Thread) {
    _aload(thread, 0)
}

type aload_1 struct {NoOperandsInstruction}
func (self *aload_1) execute(thread *rtda.Thread) {
    _aload(thread, 1)
}

type aload_2 struct {NoOperandsInstruction}
func (self *aload_2) execute(thread *rtda.Thread) {
    _aload(thread, 2)
}

type aload_3 struct {NoOperandsInstruction}
func (self *aload_3) execute(thread *rtda.Thread) {
    _aload(thread, 3)
}

func _aload(thread *rtda.Thread, index uint) {
    ref := thread.CurrentFrame().LocalVars().GetRef(index)
    thread.CurrentFrame().OperandStack().PushRef(ref)
}
