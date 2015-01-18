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

// Add float
type fadd struct {}
func (self *fadd) fetchOperands(bcr *BytecodeReader) {}
func (self *fadd) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := v1 + v2
    stack.PushFloat(result)
}

// Add int
type iadd struct {}
func (self *iadd) fetchOperands(bcr *BytecodeReader) {}
func (self *iadd) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 + v2
    stack.PushInt(result)
}

// Add long
type ladd struct {}
func (self *ladd) fetchOperands(bcr *BytecodeReader) {}
func (self *ladd) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 + v2
    stack.PushLong(result)
}
