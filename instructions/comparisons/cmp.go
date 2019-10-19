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
	if v1 > v2 {
		frame.PushInt(1)
	} else if v1 == v2 {
		frame.PushInt(0)
	} else {
		frame.PushInt(-1)
	}
}

// Compare float
type FCMP struct {
	base.NoOperandsInstruction
	g bool // long or double
}

func (instr *FCMP) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	if v1 > v2 {
		frame.PushInt(1)
	} else if v1 == v2 {
		frame.PushInt(0)
	} else if v1 < v2 {
		frame.PushInt(-1)
	} else if instr.g {
		frame.PushInt(1)
	} else {
		frame.PushInt(-1)
	}
}

// Compare double
type DCMP struct {
	base.NoOperandsInstruction
	g bool // long or double
}

func (instr *DCMP) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	if v1 > v2 {
		frame.PushInt(1)
	} else if v1 == v2 {
		frame.PushInt(0)
	} else if v1 < v2 {
		frame.PushInt(-1)
	} else if instr.g {
		frame.PushInt(1)
	} else {
		frame.PushInt(-1)
	}
}
