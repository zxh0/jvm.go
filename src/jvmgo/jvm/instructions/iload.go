package instructions

import "jvmgo/jvm/rtda"

// Load int from local variable 
type iload struct {Index8Instruction}
func (self *iload) Execute(frame *rtda.Frame) {
    _iload(frame, uint(self.index))
}

type iload_0 struct {NoOperandsInstruction}
func (self *iload_0) Execute(frame *rtda.Frame) {
    _iload(frame, 0)
}

type iload_1 struct {NoOperandsInstruction}
func (self *iload_1) Execute(frame *rtda.Frame) {
    _iload(frame, 1)
}

type iload_2 struct {NoOperandsInstruction}
func (self *iload_2) Execute(frame *rtda.Frame) {
    _iload(frame, 2)
}

type iload_3 struct {NoOperandsInstruction}
func (self *iload_3) Execute(frame *rtda.Frame) {
    _iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
    val := frame.LocalVars().GetInt(index)
    frame.OperandStack().PushInt(val)
}
