package math

import (
	"math"

	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Remainder double
type DREM struct{ base.NoOperandsInstruction }

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

// Remainder float
type FREM struct{ base.NoOperandsInstruction }

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// Remainder int
type IREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		stack.PushInt(result)
	}
}

// Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		stack.PushLong(result)
	}
}
