package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Jump subroutine
type jsr struct{ BranchInstruction }

func (self *jsr) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Jump subroutine (wide index)
type jsr_w struct {
	offset int
}

func (self *jsr_w) fetchOperands(decoder *InstructionDecoder) {
	self.offset = int(decoder.readInt32())
}
func (self *jsr_w) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Return from subroutine
type ret struct{ Index8Instruction }

func (self *ret) Execute(frame *rtda.Frame) {
	panic("todo")
}
