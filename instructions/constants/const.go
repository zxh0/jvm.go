package constants

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// xconst: Push XXX
type Const struct {
	base.NoOperandsInstruction
	K heap.Slot
	L bool
}

func (instr *Const) Execute(frame *rtda.Frame) {
	frame.Push(instr.K)
	if instr.L {
		frame.PushNull()
	}
}
