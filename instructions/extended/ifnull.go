package extended

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch if reference is null
type IfNull struct{ base.BranchInstruction }

func (instr *IfNull) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, instr.Offset)
	}
}

// Branch if reference not null
type IfNonNull struct{ base.BranchInstruction }

func (instr *IfNonNull) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, instr.Offset)
	}
}
