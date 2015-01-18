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

// Push float
type fconst_0 struct {}
func (self *fconst_0) fetchOperands(bcr *BytecodeReader) {}
func (self *fconst_0) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushFloat(0.0)
}

type fconst_1 struct {}
func (self *fconst_1) fetchOperands(bcr *BytecodeReader) {}
func (self *fconst_1) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushFloat(1.0)
}

type fconst_2 struct {}
func (self *fconst_2) fetchOperands(bcr *BytecodeReader) {}
func (self *fconst_2) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushFloat(2.0)
}
