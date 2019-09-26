package stack

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Duplicate the top operand stack value
type Dup struct{ base.NoOperandsInstruction }

func (instr *Dup) Execute(frame *rtda.Frame) {
	val := frame.Pop()
	frame.Push(val)
	frame.Push(val)
}

// Duplicate the top operand stack value and insert two values down
type DupX1 struct{ base.NoOperandsInstruction }

func (instr *DupX1) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	frame.Push(val1)
	frame.Push(val2)
	frame.Push(val1)
}

// Duplicate the top operand stack value and insert two or three values down
type DupX2 struct{ base.NoOperandsInstruction }

func (instr *DupX2) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	val3 := frame.Pop()
	frame.Push(val1)
	frame.Push(val3)
	frame.Push(val2)
	frame.Push(val1)
}

// Duplicate the top one or two operand stack values
type Dup2 struct{ base.NoOperandsInstruction }

func (instr *Dup2) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	frame.Push(val2)
	frame.Push(val1)
	frame.Push(val2)
	frame.Push(val1)
}

// Duplicate the top one or two operand stack values and insert two or three values down
type Dup2X1 struct{ base.NoOperandsInstruction }

func (instr *Dup2X1) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	val3 := frame.Pop()
	frame.Push(val2)
	frame.Push(val1)
	frame.Push(val3)
	frame.Push(val2)
	frame.Push(val1)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type Dup2X2 struct{ base.NoOperandsInstruction }

func (instr *Dup2X2) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	val3 := frame.Pop()
	val4 := frame.Pop()
	frame.Push(val2)
	frame.Push(val1)
	frame.Push(val4)
	frame.Push(val3)
	frame.Push(val2)
	frame.Push(val1)
}
