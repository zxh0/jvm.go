package instructions

import "jvmgo/rtda"

// Add double
type dadd struct {}
func (self *dadd) fetchOperands(bcr *BytecodeReader) {}
func (self *dadd) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 + v2
    stack.PushDouble(result)
}
