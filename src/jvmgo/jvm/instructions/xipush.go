package instructions

import "jvmgo/jvm/rtda"

// Push byte
type bipush struct {
    val int8
}
func (self *bipush) fetchOperands(bcr *BytecodeReader) {
    self.val = bcr.readInt8()
}
func (self *bipush) Execute(frame *rtda.Frame) {
    i := int32(self.val)
    frame.OperandStack().PushInt(i)
}

// Push short
type sipush struct {
    val int16
}
func (self *sipush) fetchOperands(bcr *BytecodeReader) {
    self.val = bcr.readInt16()
}
func (self *sipush) Execute(frame *rtda.Frame) {
    i := int32(self.val)
    frame.OperandStack().PushInt(i)
}
