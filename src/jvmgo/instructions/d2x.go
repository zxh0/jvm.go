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
