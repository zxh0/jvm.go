package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Convert long to double
type l2d struct{ NoOperandsInstruction }

func (self *l2d) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type l2f struct{ NoOperandsInstruction }

func (self *l2f) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type l2i struct{ NoOperandsInstruction }

func (self *l2i) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
