package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Compare double
type DCMPG struct{ base.NoOperandsInstruction }

func (instr *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOperandsInstruction }

func (instr *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	if v1 > v2 {
		frame.PushInt(1)
	} else if v1 == v2 {
		frame.PushInt(0)
	} else if v1 < v2 {
		frame.PushInt(-1)
	} else if gFlag {
		frame.PushInt(1)
	} else {
		frame.PushInt(-1)
	}
}
