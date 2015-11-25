package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Load int from local variable
type iload struct{ base.Index8Instruction }

func (self *iload) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

type iload_0 struct{ base.NoOperandsInstruction }

func (self *iload_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type iload_1 struct{ base.NoOperandsInstruction }

func (self *iload_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type iload_2 struct{ base.NoOperandsInstruction }

func (self *iload_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type iload_3 struct{ base.NoOperandsInstruction }

func (self *iload_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
