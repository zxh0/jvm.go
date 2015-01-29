package instructions

import "jvmgo/rtda"

// Store reference into local variable 
type astore struct {Index8Instruction}
func (self *astore) Execute(frame *rtda.Frame) {
    _astore(frame, uint(self.index))
}

type astore_0 struct {NoOperandsInstruction}
func (self *astore_0) Execute(frame *rtda.Frame) {
    _astore(frame, 0)
}

type astore_1 struct {NoOperandsInstruction}
func (self *astore_1) Execute(frame *rtda.Frame) {
    _astore(frame, 1)
}

type astore_2 struct {NoOperandsInstruction}
func (self *astore_2) Execute(frame *rtda.Frame) {
    _astore(frame, 2)
}

type astore_3 struct {NoOperandsInstruction}
func (self *astore_3) Execute(frame *rtda.Frame) {
    _astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
    ref := frame.OperandStack().PopRef()
    frame.LocalVars().SetRef(index, ref)
}
