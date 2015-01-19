package instructions

import "jvmgo/rtda"

// Branch if reference is null
type ifnull struct {BranchInstruction}
func (self *ifnull) execute(thread *rtda.Thread) {
    ref := thread.CurrentFrame().OperandStack().PopRef()
    if ref == nil {
        // todo
    }
}

// Branch if reference not null
type ifnonnull struct {BranchInstruction}
func (self *ifnonnull) execute(thread *rtda.Thread) {
    ref := thread.CurrentFrame().OperandStack().PopRef()
    if ref != nil {
        // todo
    }
}
