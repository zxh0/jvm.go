package math

import (
	"math"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Remainder double
type DRem struct{ base.NoOperandsInstruction }

func (instr *DRem) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	result := math.Mod(v1, v2) // todo
	frame.PushDouble(result)
}

// Remainder float
type FRem struct{ base.NoOperandsInstruction }

func (instr *FRem) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	frame.PushFloat(result)
}

// Remainder int
type IRem struct{ base.NoOperandsInstruction }

func (instr *IRem) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		frame.PushInt(result)
	}
}

// Remainder long
type LRem struct{ base.NoOperandsInstruction }

func (instr *LRem) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		frame.PushLong(result)
	}
}
