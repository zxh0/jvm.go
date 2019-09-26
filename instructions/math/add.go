package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Add double
type DAdd struct{ base.NoOperandsInstruction }

func (instr *DAdd) Execute(frame *rtda.Frame) {
	v1 := frame.PopDouble()
	v2 := frame.PopDouble()
	result := v1 + v2
	frame.PushDouble(result)
}

// Add float
type FAdd struct{ base.NoOperandsInstruction }

func (instr *FAdd) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	result := v1 + v2
	frame.PushFloat(result)
}

// Add int
type IAdd struct{ base.NoOperandsInstruction }

func (instr *IAdd) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	result := v1 + v2
	frame.PushInt(result)
}

// Add long
type LAdd struct{ base.NoOperandsInstruction }

func (instr *LAdd) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	result := v1 + v2
	frame.PushLong(result)
}
