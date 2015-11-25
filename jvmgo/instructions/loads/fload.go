package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Load float from local variable
type FLOAD struct{ base.Index8Instruction }

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, uint(self.Index))
}

type FLOAD_0 struct{ base.NoOperandsInstruction }

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperandsInstruction }

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperandsInstruction }

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperandsInstruction }

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
