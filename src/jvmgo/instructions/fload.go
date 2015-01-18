package instructions

import "jvmgo/rtda"

// Load float from local variable 
type fload struct {
    index uint8
}
func (self *fload) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint8()
}
func (self *fload) execute(thread *rtda.Thread) {
    _fload(thread, uint(self.index))
}

type fload_0 struct {}
func (self *fload_0) fetchOperands(bcr *BytecodeReader) {}
func (self *fload_0) execute(thread *rtda.Thread) {
    _fload(thread, 0)
}

type fload_1 struct {}
func (self *fload_1) fetchOperands(bcr *BytecodeReader) {}
func (self *fload_1) execute(thread *rtda.Thread) {
    _fload(thread, 1)
}

type fload_2 struct {}
func (self *fload_2) fetchOperands(bcr *BytecodeReader) {}
func (self *fload_2) execute(thread *rtda.Thread) {
    _fload(thread, 2)
}

type fload_3 struct {}
func (self *fload_3) fetchOperands(bcr *BytecodeReader) {}
func (self *fload_3) execute(thread *rtda.Thread) {
    _fload(thread, 3)
}

func _fload(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().LocalVars().GetFloat(index)
    thread.CurrentFrame().OperandStack().PushFloat(val)
}
