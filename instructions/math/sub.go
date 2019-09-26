package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Subtract double
type DSub struct{ base.NoOperandsInstruction }

func (instr *DSub) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	result := v1 - v2
	frame.PushDouble(result)
}

// Subtract float
type FSub struct{ base.NoOperandsInstruction }

func (instr *FSub) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	result := v1 - v2
	frame.PushFloat(result)
}

// Subtract int
type ISub struct{ base.NoOperandsInstruction }

func (instr *ISub) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	result := v1 - v2
	frame.PushInt(result)
}

// Subtract long
type LSub struct{ base.NoOperandsInstruction }

func (instr *LSub) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	result := v1 - v2
	frame.PushLong(result)
}
