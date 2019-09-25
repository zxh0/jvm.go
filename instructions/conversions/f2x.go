package conversions

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Convert float to double
type F2D struct{ base.NoOperandsInstruction }

func (instr *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct{ base.NoOperandsInstruction }

func (instr *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct{ base.NoOperandsInstruction }

func (instr *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
