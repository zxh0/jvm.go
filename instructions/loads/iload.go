package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Load int from local variable
type ILoad struct{ base.Index8Instruction }

func (instr *ILoad) Execute(frame *rtda.Frame) {
	_iload(frame, uint(instr.Index))
}

type ILoad0 struct{ base.NoOperandsInstruction }

func (instr *ILoad0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILoad1 struct{ base.NoOperandsInstruction }

func (instr *ILoad1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILoad2 struct{ base.NoOperandsInstruction }

func (instr *ILoad2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILoad3 struct{ base.NoOperandsInstruction }

func (instr *ILoad3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
