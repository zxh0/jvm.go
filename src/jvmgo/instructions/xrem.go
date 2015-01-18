package instructions

import "jvmgo/rtda"

// Remainder double
type drem struct {}
func (self *drem) fetchOperands(bcr *BytecodeReader) {}
func (self *drem) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopDouble()
    v2 := stack.PopDouble()
    result := v1 * v2 // todo
    stack.PushDouble(result)
}

// Remainder float
type frem struct {}
func (self *frem) fetchOperands(bcr *BytecodeReader) {}
func (self *frem) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := v1 * v2 // todo
    stack.PushFloat(result)
}

// Remainder int
type irem struct {}
func (self *irem) fetchOperands(bcr *BytecodeReader) {}
func (self *irem) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 % v2
    stack.PushInt(result)
}

// Remainder long
type lrem struct {}
func (self *lrem) fetchOperands(bcr *BytecodeReader) {}
func (self *lrem) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 % v2
    stack.PushLong(result)
}
