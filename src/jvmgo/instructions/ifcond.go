package instructions

import "jvmgo/rtda"

// Branch if int comparison with zero succeeds 
type ifeq struct {
    branch int16
}
func (self *ifeq) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *ifeq) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    if val == 0 {
        // todo
    }
}

type ifne struct {
    branch int16
}
func (self *ifne) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *ifne) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    if val != 0 {
        // todo
    }
}
type iflt struct {
    branch int16
}
func (self *iflt) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *iflt) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    if val < 0 {
        // todo
    }
}

type ifle struct {
    branch int16
}
func (self *ifle) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *ifle) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    if val <= 0 {
        // todo
    }
}

type ifgt struct {
    branch int16
}
func (self *ifgt) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *ifgt) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    if val > 0 {
        // todo
    }
}

type ifge struct {
    branch int16
}
func (self *ifge) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *ifge) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.PopInt()
    if val >= 0 {
        // todo
    }
}
