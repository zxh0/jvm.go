package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch if reference comparison succeeds
type IfACmpEQ struct{ base.BranchInstruction }

func (instr *IfACmpEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, instr.Offset)
	}
}

type IfACmpNE struct{ base.BranchInstruction }

func (instr *IfACmpNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, instr.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	ref2 := frame.PopRef()
	ref1 := frame.PopRef()
	return ref1 == ref2 // todo
}
