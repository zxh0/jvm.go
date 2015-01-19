package instructions

import "jvmgo/rtda"

// Convert int to byte
type i2b struct {}
func (self *i2b) fetchOperands(bcr *BytecodeReader) {}
func (self *i2b) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    b := int32(int8(i))
    stack.PushInt(b)
}
