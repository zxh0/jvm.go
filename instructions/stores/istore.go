package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Store int into local variable
type IStore struct{ base.Index8Instruction }

func (instr *IStore) Execute(frame *rtda.Frame) {
	_istore(frame, uint(instr.Index))
}

type IStore0 struct{ base.NoOperandsInstruction }

func (instr *IStore0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type IStore1 struct{ base.NoOperandsInstruction }

func (instr *IStore1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type IStore2 struct{ base.NoOperandsInstruction }

func (instr *IStore2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type IStore3 struct{ base.NoOperandsInstruction }

func (instr *IStore3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
