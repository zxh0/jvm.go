package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Branch always
type goto_ struct{ BranchInstruction }

func (self *goto_) Execute(frame *rtda.Frame) {
	branch(frame, self.offset)
}

// Branch always (wide index)
type goto_w struct {
	offset int
}

func (self *goto_w) fetchOperands(decoder *InstructionDecoder) {
	self.offset = int(decoder.readInt32())
}
func (self *goto_w) Execute(frame *rtda.Frame) {
	branch(frame, self.offset)
}
