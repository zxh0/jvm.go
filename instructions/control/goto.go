package control

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch always
type Goto struct{ base.BranchInstruction }

func (instr *Goto) Execute(frame *rtda.Frame) {
	base.Branch(frame, instr.Offset)
}

// Branch always (wide index)
type GotoW struct {
	offset int
}

func (instr *GotoW) FetchOperands(reader *base.CodeReader) {
	instr.offset = int(reader.ReadInt32())
}
func (instr *GotoW) Execute(frame *rtda.Frame) {
	base.Branch(frame, instr.offset)
}
