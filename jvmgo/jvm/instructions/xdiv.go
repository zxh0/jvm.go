package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
)

// Divide double
type ddiv struct{ NoOperandsInstruction }

func (self *ddiv) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Divide float
type fdiv struct{ NoOperandsInstruction }

func (self *fdiv) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Divide int
type idiv struct{ NoOperandsInstruction }

func (self *idiv) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 / v2
		stack.PushInt(result)
	}
}

// Divide long
type ldiv struct{ NoOperandsInstruction }

func (self *ldiv) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 / v2
		stack.PushLong(result)
	}
}
