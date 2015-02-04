package instructions

import "jvmgo/rtda"

// Branch if reference is null
type ifnull struct {BranchInstruction}
func (self *ifnull) Execute(frame *rtda.Frame) {
    ref := frame.OperandStack().PopRef()
    if ref == nil {
        branch(frame, self.offset)
    }
}

// Branch if reference not null
type ifnonnull struct {BranchInstruction}
func (self *ifnonnull) Execute(frame *rtda.Frame) {
    ref := frame.OperandStack().PopRef()
    if ref != nil {
        branch(frame, self.offset)
    }
}
