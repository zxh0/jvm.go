package stack

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Duplicate the top operand stack value
type dup struct{ base.NoOperandsInstruction }

func (self *dup) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopSlot()
	stack.PushSlot(val)
	stack.PushSlot(val)
}

// Duplicate the top operand stack value and insert two values down
type dup_x1 struct{ base.NoOperandsInstruction }

func (self *dup_x1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top operand stack value and insert two or three values down
type dup_x2 struct{ base.NoOperandsInstruction }

func (self *dup_x2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top one or two operand stack values
type dup2 struct{ base.NoOperandsInstruction }

func (self *dup2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top one or two operand stack values and insert two or three values down
type dup2_x1 struct{ base.NoOperandsInstruction }

func (self *dup2_x1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type dup2_x2 struct{ base.NoOperandsInstruction }

func (self *dup2_x2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	val4 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val4)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
