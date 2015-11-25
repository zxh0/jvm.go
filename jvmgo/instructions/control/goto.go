package control

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Branch always
type goto_ struct{ base.BranchInstruction }

func (self *goto_) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

// Branch always (wide index)
type goto_w struct {
	offset int
}

func (self *goto_w) fetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *goto_w) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
