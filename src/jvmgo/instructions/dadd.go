package instructions

import "jvmgo/rtda"

// Add double
type dadd struct {}
func (self *dadd) fetchOperands(bcr *BytecodeReader) {}

func (self *dadd) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    d1 := stack.PopDouble()
    d2 := stack.PopDouble()
    result := d1 + d2
    stack.PushDouble(result)
}
