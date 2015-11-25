package constants

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Push null
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushNull()
}

// Push double
type DCONST_0 struct{ base.NoOperandsInstruction }

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Push float
type FCONST_0 struct{ base.NoOperandsInstruction }

func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperandsInstruction }

func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Push int constant
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// Push long constant
type LCONST_0 struct{ base.NoOperandsInstruction }

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
