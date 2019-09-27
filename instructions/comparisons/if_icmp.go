package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch if int comparison succeeds
type IfICmpEQ struct{ base.BranchInstruction }

func (instr *IfICmpEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IfICmpNE struct{ base.BranchInstruction }

func (instr *IfICmpNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IfICmpLT struct{ base.BranchInstruction }

func (instr *IfICmpLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IfICmpLE struct{ base.BranchInstruction }

func (instr *IfICmpLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IfICmpGT struct{ base.BranchInstruction }

func (instr *IfICmpGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IfICmpGE struct{ base.BranchInstruction }

func (instr *IfICmpGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, instr.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	val2 = frame.PopInt()
	val1 = frame.PopInt()
	return
}
