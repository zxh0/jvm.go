package control

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewXReturn(d bool) *XReturn { return &XReturn{d: d} }

// Return void from method
type Return struct{ base.NoOperandsInstruction }

func (instr *Return) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	thread.PopFrame()
}

// xreturn: Return XXX from method
type XReturn struct {
	base.NoOperandsInstruction
	d bool
}

func (instr *XReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.PopL(instr.d)
	invokerFrame.PushL(ref, instr.d)
}
