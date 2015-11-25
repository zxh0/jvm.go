package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Negate double
type dneg struct{ NoOperandsInstruction }

func (self *dneg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type fneg struct{ NoOperandsInstruction }

func (self *fneg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type ineg struct{ NoOperandsInstruction }

func (self *ineg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type lneg struct{ NoOperandsInstruction }

func (self *lneg) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
