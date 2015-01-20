package instructions

import "jvmgo/rtda"

// Branch if int comparison with zero succeeds 
type ifeq struct {BranchInstruction}
func (self *ifeq) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val == 0 {
        _branch(thread, self.branch)
    }
}

type ifne struct {BranchInstruction}
func (self *ifne) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val != 0 {
        _branch(thread, self.branch)
    }
}

type iflt struct {BranchInstruction}
func (self *iflt) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val < 0 {
        _branch(thread, self.branch)
    }
}

type ifle struct {BranchInstruction}
func (self *ifle) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val <= 0 {
        _branch(thread, self.branch)
    }
}

type ifgt struct {BranchInstruction}
func (self *ifgt) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val > 0 {
        _branch(thread, self.branch)
    }
}

type ifge struct {BranchInstruction}
func (self *ifge) Execute(thread *rtda.Thread) {
    val := thread.CurrentFrame().OperandStack().PopInt()
    if val >= 0 {
        _branch(thread, self.branch)
    }
}

func _branch(thread *rtda.Thread, offset int) {
    pc := thread.PC()
    pc += offset
    thread.SetPC(pc)
}
