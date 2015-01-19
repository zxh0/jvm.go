package instructions

import "jvmgo/rtda"

// Branch if int comparison succeeds 
type if_icmpeq struct {
    branch int16
}
func (self *if_icmpeq) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_icmpeq) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.PopInt()
    val2 := stack.PopInt()
    if val1 == val2 {
        // todo
    }
}

type if_icmpne struct {
    branch int16
}
func (self *if_icmpne) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_icmpne) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.PopInt()
    val2 := stack.PopInt()
    if val1 != val2 {
        // todo
    }
}

type if_icmplt struct {
    branch int16
}
func (self *if_icmplt) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_icmplt) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.PopInt()
    val2 := stack.PopInt()
    if val1 < val2 {
        // todo
    }
}

type if_icmple struct {
    branch int16
}
func (self *if_icmple) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_icmple) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.PopInt()
    val2 := stack.PopInt()
    if val1 <= val2 {
        // todo
    }
}

type if_icmpgt struct {
    branch int16
}
func (self *if_icmpgt) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_icmpgt) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.PopInt()
    val2 := stack.PopInt()
    if val1 > val2 {
        // todo
    }
}

type if_icmpge struct {
    branch int16
}
func (self *if_icmpge) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_icmpge) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.PopInt()
    val2 := stack.PopInt()
    if val1 >= val2 {
        // todo
    }
}
