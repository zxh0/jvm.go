package constants

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Push null
type AConstNull struct{ base.NoOperandsInstruction }

func (instr *AConstNull) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushNull()
}

// Push double
type DConst struct {
	base.NoOperandsInstruction
	Val float64
}

func (instr *DConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(instr.Val)
}

// Push float
type FConst struct {
	base.NoOperandsInstruction
	Val float32
}

func (instr *FConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(instr.Val)
}

// Push int constant
type IConst struct {
	base.NoOperandsInstruction
	Val int32
}

func (instr *IConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(instr.Val)
}

// Push long constant
type LConst struct {
	base.NoOperandsInstruction
	Val int64
}

func (instr *LConst) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(instr.Val)
}
