package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Load long from local variable
type LLoad struct{ base.Index8Instruction }

func (instr *LLoad) Execute(frame *rtda.Frame) {
	_lload(frame, uint(instr.Index))
}

type LLoad0 struct{ base.NoOperandsInstruction }

func (instr *LLoad0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLoad1 struct{ base.NoOperandsInstruction }

func (instr *LLoad1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLoad2 struct{ base.NoOperandsInstruction }

func (instr *LLoad2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLoad3 struct{ base.NoOperandsInstruction }

func (instr *LLoad3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
