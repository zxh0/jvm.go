package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Load long from local variable
type LLOAD struct{ base.Index8Instruction }

func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(self.Index))
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
