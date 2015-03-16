package instructions

import "github.com/zxh0/jvm.go/jvmgo/jvm/rtda"

// Boolean AND int
type iand struct{ NoOperandsInstruction }

func (self *iand) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type land struct{ NoOperandsInstruction }

func (self *land) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
