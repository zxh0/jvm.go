package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Compare float
type FCMPG struct{ base.NoOperandsInstruction }

func (instr *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (instr *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
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
