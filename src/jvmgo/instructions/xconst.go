package instructions

import "jvmgo/rtda"

// Push double
type dconst_0 struct {NoOperandsInstruction}
func (self *dconst_0) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushDouble(0.0)
}

type dconst_1 struct {NoOperandsInstruction}
func (self *dconst_1) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushDouble(1.0)
}

// Push float
type fconst_0 struct {NoOperandsInstruction}
func (self *fconst_0) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushFloat(0.0)
}

type fconst_1 struct {NoOperandsInstruction}
func (self *fconst_1) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushFloat(1.0)
}

type fconst_2 struct {NoOperandsInstruction}
func (self *fconst_2) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushFloat(2.0)
}

// Push int constant 
type iconst_m1 struct {NoOperandsInstruction}
func (self *iconst_m1) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushInt(-1)
}

type iconst_0 struct {NoOperandsInstruction}
func (self *iconst_0) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushInt(0)
}

type iconst_1 struct {NoOperandsInstruction}
func (self *iconst_1) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushInt(1)
}

type iconst_2 struct {NoOperandsInstruction}
func (self *iconst_2) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushInt(2)
}

type iconst_3 struct {NoOperandsInstruction}
func (self *iconst_3) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushInt(3)
}

type iconst_4 struct {NoOperandsInstruction}
func (self *iconst_4) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushInt(4)
}

type iconst_5 struct {NoOperandsInstruction}
func (self *iconst_5) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushInt(5)
}
