package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

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
