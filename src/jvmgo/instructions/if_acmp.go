package instructions

import "jvmgo/rtda"

// Branch if reference comparison succeeds 
type if_acmpeq struct {BranchInstruction}
func (self *if_acmpeq) Execute(thread *rtda.Thread) {
    if _acmp(thread) {
        thread.IncrPC(self.offset)
    }
}

type if_acmpne struct {BranchInstruction}
func (self *if_acmpne) Execute(thread *rtda.Thread) {
    if !_acmp(thread) {
        thread.IncrPC(self.offset)
    }
}

func _acmp(thread *rtda.Thread) (bool) {
    stack := thread.CurrentFrame().OperandStack()
    ref1 := stack.PopRef()
    ref2 := stack.PopRef()
    return ref1 == ref2 // todo
}
