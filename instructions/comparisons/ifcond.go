package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

func (instr *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (instr *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (instr *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (instr *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (instr *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (instr *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, instr.Offset)
	}
}
