package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Store long into local variable
type LStore struct{ base.Index8Instruction }

func (instr *LStore) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(instr.Index))
}

type LStore0 struct{ base.NoOperandsInstruction }

func (instr *LStore0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LStore1 struct{ base.NoOperandsInstruction }

func (instr *LStore1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LStore2 struct{ base.NoOperandsInstruction }

func (instr *LStore2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LStore3 struct{ base.NoOperandsInstruction }

func (instr *LStore3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
