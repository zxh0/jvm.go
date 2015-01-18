package instructions

import "jvmgo/rtda"

// Load double from local variable 
type dload struct {
    index uint8
}
func (self *dload) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint8()
}
func (self *dload) execute(thread *rtda.Thread) {
    _dload(thread, uint(self.index))
}

type dload_0 struct {}
func (self *dload_0) fetchOperands(bcr *BytecodeReader) {}
func (self *dload_0) execute(thread *rtda.Thread) {
    _dload(thread, 0)
}

type dload_1 struct {}
func (self *dload_1) fetchOperands(bcr *BytecodeReader) {}
func (self *dload_1) execute(thread *rtda.Thread) {
    _dload(thread, 1)
}

type dload_2 struct {}
func (self *dload_2) fetchOperands(bcr *BytecodeReader) {}
func (self *dload_2) execute(thread *rtda.Thread) {
    _dload(thread, 2)
}

type dload_3 struct {}
func (self *dload_3) fetchOperands(bcr *BytecodeReader) {}
func (self *dload_3) execute(thread *rtda.Thread) {
    _dload(thread, 3)
}

func _dload(thread *rtda.Thread, index uint) {
    val := thread.CurrentFrame().LocalVars().GetDouble(index)
    thread.CurrentFrame().OperandStack().PushDouble(val)
}
