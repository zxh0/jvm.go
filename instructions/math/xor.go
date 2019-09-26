package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Boolean XOR int
type IXor struct{ base.NoOperandsInstruction }

func (instr *IXor) Execute(frame *rtda.Frame) {
	v1 := frame.PopInt()
	v2 := frame.PopInt()
	result := v1 ^ v2
	frame.PushInt(result)
}

// Boolean XOR long
type LXor struct{ base.NoOperandsInstruction }

func (instr *LXor) Execute(frame *rtda.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	result := v1 ^ v2
	frame.PushLong(result)
}
