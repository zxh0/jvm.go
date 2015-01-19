package instructions

import "jvmgo/rtda"

// Branch if reference is null
type ifnull struct {
    branch int16
}
func (self *ifnull) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *ifnull) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    ref := stack.PopRef()
    if ref == nil {
        // todo
    }
}

// Branch if reference not null
type ifnonnull struct {
    branch int16
}
func (self *ifnonnull) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *ifnonnull) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    ref := stack.PopRef()
    if ref != nil {
        // todo
    }
}
