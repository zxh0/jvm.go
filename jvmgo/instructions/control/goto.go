package control

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
