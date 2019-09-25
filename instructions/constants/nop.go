package constants

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (instr *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
