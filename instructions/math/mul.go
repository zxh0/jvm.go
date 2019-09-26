package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Multiply double
type DMul struct{ base.NoOperandsInstruction }

func (instr *DMul) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	result := v1 * v2
	frame.PushDouble(result)
}

// Multiply float
type FMul struct{ base.NoOperandsInstruction }

func (instr *FMul) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	result := v1 * v2
	frame.PushFloat(result)
}

// Multiply int
type IMul struct{ base.NoOperandsInstruction }

func (instr *IMul) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	result := v1 * v2
	frame.PushInt(result)
}

// Multiply long
type LMul struct{ base.NoOperandsInstruction }

func (instr *LMul) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	result := v1 * v2
	frame.PushLong(result)
}
