package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Boolean OR int
type IOr struct{ base.NoOperandsInstruction }

func (instr *IOr) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	result := v1 | v2
	frame.PushInt(result)
}

// Boolean OR long
type LOr struct{ base.NoOperandsInstruction }

func (instr *LOr) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	result := v1 | v2
	frame.PushLong(result)
}
