package conversions

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Convert int to byte
type I2B struct{ base.NoOperandsInstruction }

func (instr *I2B) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	b := int32(int8(i))
	frame.PushInt(b)
}

// Convert int to char
type I2C struct{ base.NoOperandsInstruction }

func (instr *I2C) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	c := int32(uint16(i))
	frame.PushInt(c)
}

// Convert int to short
type I2S struct{ base.NoOperandsInstruction }

func (instr *I2S) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	s := int32(int16(i))
	frame.PushInt(s)
}

// Convert int to long
type I2L struct{ base.NoOperandsInstruction }

func (instr *I2L) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	l := int64(i)
	frame.PushLong(l)
}

// Convert int to float
type I2F struct{ base.NoOperandsInstruction }

func (instr *I2F) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	f := float32(i)
	frame.PushFloat(f)
}

// Convert int to double
type I2D struct{ base.NoOperandsInstruction }

func (instr *I2D) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	d := float64(i)
	frame.PushDouble(d)
}
