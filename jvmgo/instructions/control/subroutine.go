package control

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Jump subroutine
type JSR struct{ base.BranchInstruction }

func (self *JSR) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Jump subroutine (wide index)
type JSR_W struct {
	offset int
}

func (self *JSR_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *JSR_W) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Return from subroutine
type RET struct{ base.Index8Instruction }

func (self *RET) Execute(frame *rtda.Frame) {
	panic("todo")
}
