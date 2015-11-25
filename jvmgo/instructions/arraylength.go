package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Get length of array
type arraylength struct{ NoOperandsInstruction }

func (self *arraylength) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()

	if arrRef == nil {
		frame.Thread().ThrowNPE()
		return
	}

	arrLen := rtc.ArrayLength(arrRef)
	stack.PushInt(arrLen)
}
