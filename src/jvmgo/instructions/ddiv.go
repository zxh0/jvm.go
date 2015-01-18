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

// Divide float
type fdiv struct {}
func (self *fdiv) fetchOperands(bcr *BytecodeReader) {}
func (self *fdiv) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopFloat()
    v2 := stack.PopFloat()
    result := v1 / v2
    stack.PushFloat(result)
}

// Divide int
type idiv struct {}
func (self *idiv) fetchOperands(bcr *BytecodeReader) {}
func (self *idiv) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 / v2
    stack.PushInt(result)
}

// Divide long
type ldiv struct {}
func (self *ldiv) fetchOperands(bcr *BytecodeReader) {}
func (self *ldiv) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopLong()
    v2 := stack.PopLong()
    result := v1 / v2
    stack.PushLong(result)
}
