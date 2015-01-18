package instructions

import "jvmgo/rtda"

// Divide double
type ddiv struct {}
func (self *ddiv) fetchOperands(bcr *BytecodeReader) {}
func (self *ddiv) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 / v2
    stack.PushDouble(result)
}
