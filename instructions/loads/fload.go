package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Load float from local variable
type FLoad struct{ base.Index8Instruction }

func (instr *FLoad) Execute(frame *rtda.Frame) {
	_fload(frame, uint(instr.Index))
}

type FLoad0 struct{ base.NoOperandsInstruction }

func (instr *FLoad0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLoad1 struct{ base.NoOperandsInstruction }

func (instr *FLoad1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLoad2 struct{ base.NoOperandsInstruction }

func (instr *FLoad2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLoad3 struct{ base.NoOperandsInstruction }

func (instr *FLoad3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
