package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Boolean AND int
type IAnd struct{ base.NoOperandsInstruction }

func (instr *IAnd) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	result := v1 & v2
	frame.PushInt(result)
}

// Boolean AND long
type LAnd struct{ base.NoOperandsInstruction }

func (instr *LAnd) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	result := v1 & v2
	frame.PushLong(result)
}
