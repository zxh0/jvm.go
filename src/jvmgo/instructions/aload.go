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

type aload_0 struct {}
func (self *aload_0) fetchOperands(bcr *BytecodeReader) {}
func (self *aload_0) execute(thread *rtda.Thread) {
    _aload(thread, 0)
}

func _aload(thread *rtda.Thread, index uint) {
    ref := thread.CurrentFrame().LocalVars().GetRef(index)
    thread.CurrentFrame().OperandStack().PushRef(ref)
}
