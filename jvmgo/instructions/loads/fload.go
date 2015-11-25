package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Load float from local variable
type fload struct{ base.Index8Instruction }

func (self *fload) Execute(frame *rtda.Frame) {
	_fload(frame, uint(self.Index))
}

type fload_0 struct{ base.NoOperandsInstruction }

func (self *fload_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type fload_1 struct{ base.NoOperandsInstruction }

func (self *fload_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type fload_2 struct{ base.NoOperandsInstruction }

func (self *fload_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type fload_3 struct{ base.NoOperandsInstruction }

func (self *fload_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
