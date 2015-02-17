package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
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
