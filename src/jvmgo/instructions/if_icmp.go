package instructions

import "jvmgo/rtda"

// Branch if int comparison succeeds 
type if_icmpeq struct {BranchInstruction}
func (self *if_icmpeq) Execute(thread *rtda.Thread) {
    if val1, val2 := pop2Ints(thread); val1 == val2 {
        thread.IncrPC(self.offset)
    }
}

type if_icmpne struct {BranchInstruction}
func (self *if_icmpne) Execute(thread *rtda.Thread) {
    if val1, val2 := pop2Ints(thread); val1 != val2 {
        thread.IncrPC(self.offset)
    }
}

type if_icmplt struct {BranchInstruction}
func (self *if_icmplt) Execute(thread *rtda.Thread) {
    if val1, val2 := pop2Ints(thread); val1 < val2 {
        thread.IncrPC(self.offset)
    }
}

type if_icmple struct {BranchInstruction}
func (self *if_icmple) Execute(thread *rtda.Thread) {
    if val1, val2 := pop2Ints(thread); val1 <= val2 {
        thread.IncrPC(self.offset)
    }
}

type if_icmpgt struct {BranchInstruction}
func (self *if_icmpgt) Execute(thread *rtda.Thread) {
    if val1, val2 := pop2Ints(thread); val1 > val2 {
        thread.IncrPC(self.offset)
    }
}

type if_icmpge struct {BranchInstruction}
func (self *if_icmpge) Execute(thread *rtda.Thread) {
    if val1, val2 := pop2Ints(thread); val1 >= val2 {
        thread.IncrPC(self.offset)
    }
}

func pop2Ints(thread *rtda.Thread) (val1, val2 int32) {
    stack := thread.CurrentFrame().OperandStack()
    val1 = stack.PopInt()
    val2 = stack.PopInt()
    return
}
