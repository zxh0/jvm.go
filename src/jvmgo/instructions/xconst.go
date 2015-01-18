package instructions

import "jvmgo/rtda"

// Push double
type dconst_0 struct {}
func (self *dconst_0) fetchOperands(bcr *BytecodeReader) {}
func (self *dconst_0) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushDouble(0.0)
}

type dconst_1 struct {}
func (self *dconst_1) fetchOperands(bcr *BytecodeReader) {}
func (self *dconst_1) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushDouble(1.0)
}
