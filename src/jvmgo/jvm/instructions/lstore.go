package instructions

import "jvmgo/jvm/rtda"

// Store long into local variable 
type lstore struct {Index8Instruction}
func (self *lstore) Execute(frame *rtda.Frame) {
    _lstore(frame, uint(self.index))
}

type lstore_0 struct {NoOperandsInstruction}
func (self *lstore_0) Execute(frame *rtda.Frame) {
    _lstore(frame, 0)
}

type lstore_1 struct {NoOperandsInstruction}
func (self *lstore_1) Execute(frame *rtda.Frame) {
    _lstore(frame, 1)
}

type lstore_2 struct {NoOperandsInstruction}
func (self *lstore_2) Execute(frame *rtda.Frame) {
    _lstore(frame, 2)
}

type lstore_3 struct {NoOperandsInstruction}
func (self *lstore_3) Execute(frame *rtda.Frame) {
    _lstore(frame, 3)
}

func _lstore(frame *rtda.Frame, index uint) {
    val := frame.OperandStack().PopLong()
    frame.LocalVars().SetLong(index, val)
}
