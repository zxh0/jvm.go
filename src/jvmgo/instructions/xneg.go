package instructions

import "jvmgo/rtda"

// Negate double
type dneg struct {}
func (self *dneg) fetchOperands(bcr *BytecodeReader) {}
func (self *dneg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopDouble()
    stack.PushDouble(-val)
}

// Negate float
type fneg struct {}
func (self *fneg) fetchOperands(bcr *BytecodeReader) {}
func (self *fneg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopFloat()
    stack.PushFloat(-val)
}

// Negate int
type ineg struct {}
func (self *ineg) fetchOperands(bcr *BytecodeReader) {}
func (self *ineg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    stack.PushInt(-val)
}

// Negate long
type lneg struct {}
func (self *lneg) fetchOperands(bcr *BytecodeReader) {}
func (self *lneg) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopLong()
    stack.PushLong(-val)
}
