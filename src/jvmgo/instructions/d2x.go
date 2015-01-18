package instructions

import "jvmgo/rtda"

// Convert double to float
type d2f struct {}
func (self *d2f) fetchOperands(bcr *BytecodeReader) {}
func (self *d2f) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    d := stack.PopDouble()
    f := float32(d)
    stack.PushFloat(f)
}

// Convert double to int
type d2i struct {}
func (self *d2i) fetchOperands(bcr *BytecodeReader) {}
func (self *d2i) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    d := stack.PopDouble()
    i := int32(d)
    stack.PushInt(i)
}
