package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Store int into local variable
type IStore struct{ base.Index8Instruction }

func (instr *IStore) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(instr.Index, val)
}

// Store long into local variable
type LStore struct{ base.Index8Instruction }

func (instr *LStore) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(instr.Index, val)
}

// Store float into local variable
type FStore struct{ base.Index8Instruction }

func (instr *FStore) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(instr.Index, val)
}

// Store double into local variable
type DStore struct{ base.Index8Instruction }

func (instr *DStore) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(instr.Index, val)
}

// Store reference into local variable
type AStore struct{ base.Index8Instruction }

func (instr *AStore) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(instr.Index, ref)
}
