package instructions

import "jvmgo/rtda"

// Subtract double
type dsub struct {}
func (self *dsub) fetchOperands(bcr *BytecodeReader) {}
func (self *dsub) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 - v2
    stack.PushDouble(result)
}

// Subtract float
type fsub struct {}
func (self *fsub) fetchOperands(bcr *BytecodeReader) {}
func (self *fsub) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := v1 - v2
    stack.PushFloat(result)
}

// Subtract int
type isub struct {}
func (self *isub) fetchOperands(bcr *BytecodeReader) {}
func (self *isub) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 - v2
    stack.PushInt(result)
}

// Subtract long
type lsub struct {}
func (self *lsub) fetchOperands(bcr *BytecodeReader) {}
func (self *lsub) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 - v2
    stack.PushLong(result)
}
