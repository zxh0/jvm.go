package control

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Jump subroutine
type jsr struct{ base.BranchInstruction }

func (self *jsr) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Jump subroutine (wide index)
type jsr_w struct {
	offset int
}

func (self *jsr_w) fetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *jsr_w) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Return from subroutine
type ret struct{ base.Index8Instruction }

func (self *ret) Execute(frame *rtda.Frame) {
	panic("todo")
}
