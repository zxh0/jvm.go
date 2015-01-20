package instructions

import "jvmgo/rtda"

// Branch if int comparison with zero succeeds 
type ifeq struct {BranchInstruction}
func (self *ifeq) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val == 0 {
        // todo
    }
}

type ifne struct {BranchInstruction}
func (self *ifne) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val != 0 {
        // todo
    }
}

type iflt struct {BranchInstruction}
func (self *iflt) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val < 0 {
        // todo
    }
}

type ifle struct {BranchInstruction}
func (self *ifle) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val <= 0 {
        // todo
    }
}

type ifgt struct {BranchInstruction}
func (self *ifgt) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val > 0 {
        // todo
    }
}

type ifge struct {BranchInstruction}
func (self *ifge) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val >= 0 {
        // todo
    }
}
