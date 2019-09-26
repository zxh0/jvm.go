package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Store reference into local variable
type AStore struct{ base.Index8Instruction }

func (instr *AStore) Execute(frame *rtda.Frame) {
	_astore(frame, uint(instr.Index))
}

type AStore0 struct{ base.NoOperandsInstruction }

func (instr *AStore0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type AStore1 struct{ base.NoOperandsInstruction }

func (instr *AStore1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type AStore2 struct{ base.NoOperandsInstruction }

func (instr *AStore2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type AStore3 struct{ base.NoOperandsInstruction }

func (instr *AStore3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
