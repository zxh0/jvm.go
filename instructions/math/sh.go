package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Shift left int
type ISHL struct{ base.NoOperandsInstruction }

func (instr *ISHL) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	frame.PushInt(result)
}

// Arithmetic shift right int
type ISHR struct{ base.NoOperandsInstruction }

func (instr *ISHR) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	frame.PushInt(result)
}

// Logical shift right int
type IUSHR struct{ base.NoOperandsInstruction }

func (instr *IUSHR) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	frame.PushInt(result)
}

// Shift left long
type LSHL struct{ base.NoOperandsInstruction }

func (instr *LSHL) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	frame.PushLong(result)
}

// Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

func (instr *LSHR) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	frame.PushLong(result)
}

// Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }

func (instr *LUSHR) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	frame.PushLong(result)
}
