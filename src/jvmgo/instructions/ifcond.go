package instructions

import "jvmgo/rtda"

// Branch if int comparison with zero succeeds 
type ifeq struct {BranchInstruction}
func (self *ifeq) execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val == 0 {
        // todo
    }
}

type ifne struct {BranchInstruction}
func (self *ifne) execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val != 0 {
        // todo
    }
}

type iflt struct {BranchInstruction}
func (self *iflt) execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val < 0 {
        // todo
    }
}

type ifle struct {BranchInstruction}
func (self *ifle) execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val <= 0 {
        // todo
    }
}

type ifgt struct {BranchInstruction}
func (self *ifgt) execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val > 0 {
        // todo
    }
}

type ifge struct {BranchInstruction}
func (self *ifge) execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val >= 0 {
        // todo
    }
}
