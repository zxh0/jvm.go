package instructions

import "jvmgo/rtda"

// Branch if reference comparison succeeds 
type if_acmpeq struct {BranchInstruction}
func (self *if_acmpeq) execute(thread *rtda.Thread) {
    if _acmp(thread) {
        // todo
    }
}

type if_acmpne struct {BranchInstruction}
func (self *if_acmpne) execute(thread *rtda.Thread) {
    if !_acmp(thread) {
        // todo
    }
}

func _acmp(thread *rtda.Thread) (bool) {
    stack := thread.CurrentFrame().OperandStack()
    ref1 := stack.PopRef()
    ref2 := stack.PopRef()
    return ref1 == ref2 // todo
}
