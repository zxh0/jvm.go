package instructions

import "jvmgo/rtda"

// Branch if int comparison succeeds 
type if_icmpeq struct {BranchInstruction}
func (self *if_icmpeq) execute(thread *rtda.Thread) {
    if val1, val2 := popInts(thread); val1 == val2 {
        // todo
    }
}

type if_icmpne struct {BranchInstruction}
func (self *if_icmpne) execute(thread *rtda.Thread) {
    if val1, val2 := popInts(thread); val1 != val2 {
        // todo
    }
}

type if_icmplt struct {BranchInstruction}
func (self *if_icmplt) execute(thread *rtda.Thread) {
    if val1, val2 := popInts(thread); val1 < val2 {
        // todo
    }
}

type if_icmple struct {BranchInstruction}
func (self *if_icmple) execute(thread *rtda.Thread) {
    if val1, val2 := popInts(thread); val1 <= val2 {
        // todo
    }
}

type if_icmpgt struct {BranchInstruction}
func (self *if_icmpgt) execute(thread *rtda.Thread) {
    if val1, val2 := popInts(thread); val1 > val2 {
        // todo
    }
}

type if_icmpge struct {BranchInstruction}
func (self *if_icmpge) execute(thread *rtda.Thread) {
    if val1, val2 := popInts(thread); val1 >= val2 {
        // todo
    }
}

func popInts(thread *rtda.Thread) (val1, val2 int32) {
    stack := thread.CurrentFrame().OperandStack()
    val1 = stack.PopInt()
    val2 = stack.PopInt()
    return
}
