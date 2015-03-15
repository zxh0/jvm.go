package instructions

import "jvmgo/jvm/rtda"

// Load long from local variable
type lload struct{ Index8Instruction }

func (self *lload) Execute(frame *rtda.Frame) {
	_lload(frame, uint(self.index))
}

type lload_0 struct{ NoOperandsInstruction }

func (self *lload_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type lload_1 struct{ NoOperandsInstruction }

func (self *lload_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type lload_2 struct{ NoOperandsInstruction }

func (self *lload_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type lload_3 struct{ NoOperandsInstruction }

func (self *lload_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
