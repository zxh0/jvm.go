package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Divide double
type DDiv struct{ base.NoOperandsInstruction }

func (instr *DDiv) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	result := v1 / v2
	frame.PushDouble(result)
}

// Divide float
type FDiv struct{ base.NoOperandsInstruction }

func (instr *FDiv) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	result := v1 / v2
	frame.PushFloat(result)
}

// Divide int
type IDiv struct{ base.NoOperandsInstruction }

func (instr *IDiv) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()

	if v2 == 0 {
		frame.Thread.ThrowDivByZero()
	} else {
		result := v1 / v2
		frame.PushInt(result)
	}
}

// Divide long
type LDiv struct{ base.NoOperandsInstruction }

func (instr *LDiv) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()

	if v2 == 0 {
		frame.Thread.ThrowDivByZero()
	} else {
		result := v1 / v2
		frame.PushLong(result)
	}
}
