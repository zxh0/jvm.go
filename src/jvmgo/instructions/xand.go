package instructions

import "jvmgo/rtda"

// Boolean AND int
type iand struct {}
func (self *iand) fetchOperands(bcr *BytecodeReader) {}
func (self *iand) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    v1 := stack.PopInt()
    v2 := stack.PopInt()
    result := v1 & v2
    stack.PushInt(result)
}
