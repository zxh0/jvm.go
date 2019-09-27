package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch if int comparison with zero succeeds
type IfEQ struct{ base.BranchInstruction }

func (instr *IfEQ) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val == 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IfNE struct{ base.BranchInstruction }

func (instr *IfNE) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val != 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IfLT struct{ base.BranchInstruction }

func (instr *IfLT) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val < 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IfLE struct{ base.BranchInstruction }

func (instr *IfLE) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val <= 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IfGT struct{ base.BranchInstruction }

func (instr *IfGT) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val > 0 {
		base.Branch(frame, instr.Offset)
	}
}

type IfGE struct{ base.BranchInstruction }

func (instr *IfGE) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val >= 0 {
		base.Branch(frame, instr.Offset)
	}
}
