package instructions

import "jvmgo/rtda"

// Multiply double
type dmul struct {}
func (self *dmul) fetchOperands(bcr *BytecodeReader) {}
func (self *dmul) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 * v2
    stack.PushDouble(result)
}

// Multiply float
type fmul struct {}
func (self *fmul) fetchOperands(bcr *BytecodeReader) {}
func (self *fmul) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := v1 * v2
    stack.PushFloat(result)
}

// Multiply int
type imul struct {}
func (self *imul) fetchOperands(bcr *BytecodeReader) {}
func (self *imul) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 * v2
    stack.PushInt(result)
}

// Multiply long
type lmul struct {}
func (self *lmul) fetchOperands(bcr *BytecodeReader) {}
func (self *lmul) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 * v2
    stack.PushLong(result)
}
