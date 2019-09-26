package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }

func (instr *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (instr *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (instr *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (instr *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (instr *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, instr.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (instr *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, instr.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	val2 = frame.PopInt()
	val1 = frame.PopInt()
	return
}
