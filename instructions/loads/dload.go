package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Load double from local variable
type DLoad struct{ base.Index8Instruction }

func (instr *DLoad) Execute(frame *rtda.Frame) {
	_dload(frame, uint(instr.Index))
}

type DLoad0 struct{ base.NoOperandsInstruction }

func (instr *DLoad0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLoad1 struct{ base.NoOperandsInstruction }

func (instr *DLoad1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLoad2 struct{ base.NoOperandsInstruction }

func (instr *DLoad2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLoad3 struct{ base.NoOperandsInstruction }

func (instr *DLoad3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
