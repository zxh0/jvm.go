package instructions

import "jvmgo/rtda"

// Push byte
type bipush struct {
    val int8
}
func (self *bipush) fetchOperands(bcr *BytecodeReader) {
    self.val = bcr.readInt8()
}
func (self *bipush) Execute(thread *rtda.Thread) {
    i := int32(self.val)
    thread.CurrentFrame().OperandStack().PushInt(i)
}

// Push short
type sipush struct {
    val int16
}
func (self *sipush) fetchOperands(bcr *BytecodeReader) {
    self.val = bcr.readInt16()
}
func (self *sipush) Execute(thread *rtda.Thread) {
    i := int32(self.val)
    thread.CurrentFrame().OperandStack().PushInt(i)
}
