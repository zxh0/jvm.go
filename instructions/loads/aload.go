package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Load reference from local variable
type ALoad struct{ base.Index8Instruction }

func (instr *ALoad) Execute(frame *rtda.Frame) {
	_aload(frame, uint(instr.Index))
}

type ALoad0 struct{ base.NoOperandsInstruction }

func (instr *ALoad0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALoad1 struct{ base.NoOperandsInstruction }

func (instr *ALoad1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALoad2 struct{ base.NoOperandsInstruction }

func (instr *ALoad2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALoad3 struct{ base.NoOperandsInstruction }

func (instr *ALoad3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
