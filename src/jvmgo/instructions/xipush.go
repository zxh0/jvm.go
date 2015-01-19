package instructions

import "jvmgo/rtda"

// Push byte
type bipush struct {
    _byte int8
}
func (self *bipush) fetchOperands(bcr *BytecodeReader) {
    self._byte = bcr.readInt8()
}
func (self *bipush) execute(thread *rtda.Thread) {
    i := int32(self._byte)
    thread.CurrentFrame().OperandStack().PushInt(i)
}
