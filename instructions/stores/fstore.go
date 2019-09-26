package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Store float into local variable
type FStore struct{ base.Index8Instruction }

func (instr *FStore) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(instr.Index))
}

type FStore0 struct{ base.NoOperandsInstruction }

func (instr *FStore0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FStore1 struct{ base.NoOperandsInstruction }

func (instr *FStore1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FStore2 struct{ base.NoOperandsInstruction }

func (instr *FStore2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FStore3 struct{ base.NoOperandsInstruction }

func (instr *FStore3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
