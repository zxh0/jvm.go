package conversions

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Convert double to float
type D2F struct{ base.NoOperandsInstruction }

func (instr *D2F) Execute(frame *rtda.Frame) {
	d := frame.PopDouble()
	f := float32(d)
	frame.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (instr *D2I) Execute(frame *rtda.Frame) {
	d := frame.PopDouble()
	i := int32(d)
	frame.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (instr *D2L) Execute(frame *rtda.Frame) {
	d := frame.PopDouble()
	l := int64(d)
	frame.PushLong(l)
}
