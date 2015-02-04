package instructions

import "jvmgo/jvm/rtda"

// Branch if reference comparison succeeds 
type if_acmpeq struct {BranchInstruction}
func (self *if_acmpeq) Execute(frame *rtda.Frame) {
    if _acmp(frame) {
        branch(frame, self.offset)
    }
}

type if_acmpne struct {BranchInstruction}
func (self *if_acmpne) Execute(frame *rtda.Frame) {
    if !_acmp(frame) {
        branch(frame, self.offset)
    }
}

func _acmp(frame *rtda.Frame) (bool) {
    stack := frame.OperandStack()
    ref2 := stack.PopRef()
    ref1 := stack.PopRef()
    return ref1 == ref2 // todo
}
