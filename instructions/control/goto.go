package control

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (instr *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, instr.Offset)
}

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (instr *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	instr.offset = int(reader.ReadInt32())
}
func (instr *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, instr.offset)
}
