package control

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Jump subroutine
type JSR struct{ base.BranchInstruction }

func (instr *JSR) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Jump subroutine (wide index)
type JSR_W struct {
	offset int
}

func (instr *JSR_W) FetchOperands(reader *base.CodeReader) {
	instr.offset = int(reader.ReadInt32())
}
func (instr *JSR_W) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Return from subroutine
type RET struct{ base.Index8Instruction }

func (instr *RET) Execute(frame *rtda.Frame) {
	panic("todo")
}
