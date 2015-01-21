package instructions

import "jvmgo/rtda"

// Branch if reference is null
type ifnull struct {BranchInstruction}
func (self *ifnull) Execute(thread *rtda.Thread) {
    ref := thread.CurrentFrame().OperandStack().PopRef()
    if ref == nil {
        branch(thread, self.offset)
    }
}

// Branch if reference not null
type ifnonnull struct {BranchInstruction}
func (self *ifnonnull) Execute(thread *rtda.Thread) {
    ref := thread.CurrentFrame().OperandStack().PopRef()
    if ref != nil {
        branch(thread, self.offset)
    }
}
