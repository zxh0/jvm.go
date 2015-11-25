package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
)

// Load double from local variable
type dload struct{ Index8Instruction }

func (self *dload) Execute(frame *rtda.Frame) {
	_dload(frame, uint(self.index))
}

type dload_0 struct{ NoOperandsInstruction }

func (self *dload_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type dload_1 struct{ NoOperandsInstruction }

func (self *dload_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type dload_2 struct{ NoOperandsInstruction }

func (self *dload_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type dload_3 struct{ NoOperandsInstruction }

func (self *dload_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
