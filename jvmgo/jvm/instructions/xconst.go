package instructions

import "github.com/zxh0/jvm.go/jvmgo/jvm/rtda"

// Push null
type aconst_null struct{ NoOperandsInstruction }

func (self *aconst_null) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushNull()
}

// Push double
type dconst_0 struct{ NoOperandsInstruction }

func (self *dconst_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type dconst_1 struct{ NoOperandsInstruction }

func (self *dconst_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Push float
type fconst_0 struct{ NoOperandsInstruction }

func (self *fconst_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type fconst_1 struct{ NoOperandsInstruction }

func (self *fconst_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type fconst_2 struct{ NoOperandsInstruction }

func (self *fconst_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Push int constant
type iconst_m1 struct{ NoOperandsInstruction }

func (self *iconst_m1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type iconst_0 struct{ NoOperandsInstruction }

func (self *iconst_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type iconst_1 struct{ NoOperandsInstruction }

func (self *iconst_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type iconst_2 struct{ NoOperandsInstruction }

func (self *iconst_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type iconst_3 struct{ NoOperandsInstruction }

func (self *iconst_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type iconst_4 struct{ NoOperandsInstruction }

func (self *iconst_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type iconst_5 struct{ NoOperandsInstruction }

func (self *iconst_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// Push long constant
type lconst_0 struct{ NoOperandsInstruction }

func (self *lconst_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type lconst_1 struct{ NoOperandsInstruction }

func (self *lconst_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
