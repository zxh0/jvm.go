package conversions

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Convert long to double
type L2D struct{ base.NoOperandsInstruction }

func (instr *L2D) Execute(frame *rtda.Frame) {
	l := frame.PopLong()
	d := float64(l)
	frame.PushDouble(d)
}

// Convert long to float
type L2F struct{ base.NoOperandsInstruction }

func (instr *L2F) Execute(frame *rtda.Frame) {
	l := frame.PopLong()
	f := float32(l)
	frame.PushFloat(f)
}

// Convert long to int
type L2I struct{ base.NoOperandsInstruction }

func (instr *L2I) Execute(frame *rtda.Frame) {
	l := frame.PopLong()
	i := int32(l)
	frame.PushInt(i)
}
