package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Negate double
type DNeg struct{ base.NoOperandsInstruction }

func (instr *DNeg) Execute(frame *rtda.Frame) {
	val := frame.PopDouble()
	frame.PushDouble(-val)
}

// Negate float
type FNeg struct{ base.NoOperandsInstruction }

func (instr *FNeg) Execute(frame *rtda.Frame) {
	val := frame.PopFloat()
	frame.PushFloat(-val)
}

// Negate int
type INeg struct{ base.NoOperandsInstruction }

func (instr *INeg) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	frame.PushInt(-val)
}

// Negate long
type LNeg struct{ base.NoOperandsInstruction }

func (instr *LNeg) Execute(frame *rtda.Frame) {
	val := frame.PopLong()
	frame.PushLong(-val)
}
