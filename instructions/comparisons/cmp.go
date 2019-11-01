package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewLCMP() *LCMP  { return &LCMP{} }
func NewFCMPG() *FCMP { return &FCMP{g: true} }
func NewFCMPL() *FCMP { return &FCMP{g: false} }
func NewDCMPG() *DCMP { return &DCMP{g: true} }
func NewDCMPL() *DCMP { return &DCMP{g: false} }

// Compare long
type LCMP struct{ base.NoOperandsInstruction }

func (instr *LCMP) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	switch {
	case v1 > v2:
		frame.PushInt(1)
	case v1 == v2:
		frame.PushInt(0)
	default:
		frame.PushInt(-1)
	}
}

// Compare float
type FCMP struct {
	base.NoOperandsInstruction
	g bool
}

func (instr *FCMP) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	switch {
	case v1 > v2:
		frame.PushInt(1)
	case v1 == v2:
		frame.PushInt(0)
	case v1 < v2:
		frame.PushInt(-1)
	case instr.g:
		frame.PushInt(1)
	default:
		frame.PushInt(-1)
	}
}

// Compare double
type DCMP struct {
	base.NoOperandsInstruction
	g bool
}

func (instr *DCMP) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	switch {
	case v1 > v2:
		frame.PushInt(1)
	case v1 == v2:
		frame.PushInt(0)
	case v1 < v2:
		frame.PushInt(-1)
	case instr.g:
		frame.PushInt(1)
	default:
		frame.PushInt(-1)
	}
}
