package instructions

import "jvmgo/jvm/rtda"

// Store double into local variable 
type dstore struct {Index8Instruction}
func (self *dstore) Execute(frame *rtda.Frame) {
    _dstore(frame, uint(self.index))
}

type dstore_0 struct {NoOperandsInstruction}
func (self *dstore_0) Execute(frame *rtda.Frame) {
    _dstore(frame, 0)
}

type dstore_1 struct {NoOperandsInstruction}
func (self *dstore_1) Execute(frame *rtda.Frame) {
    _dstore(frame, 1)
}

type dstore_2 struct {NoOperandsInstruction}
func (self *dstore_2) Execute(frame *rtda.Frame) {
    _dstore(frame, 2)
}

type dstore_3 struct {NoOperandsInstruction}
func (self *dstore_3) Execute(frame *rtda.Frame) {
    _dstore(frame, 3)
}

func _dstore(frame *rtda.Frame, index uint) {
    val := frame.OperandStack().PopDouble()
    frame.LocalVars().SetDouble(index, val)
}
