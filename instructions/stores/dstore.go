package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Store double into local variable
type DStore struct{ base.Index8Instruction }

func (instr *DStore) Execute(frame *rtda.Frame) {
	_dstore(frame, uint(instr.Index))
}

type DStore0 struct{ base.NoOperandsInstruction }

func (instr *DStore0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DStore1 struct{ base.NoOperandsInstruction }

func (instr *DStore1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DStore2 struct{ base.NoOperandsInstruction }

func (instr *DStore2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DStore3 struct{ base.NoOperandsInstruction }

func (instr *DStore3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
