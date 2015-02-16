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
	arrLen := rtc.ArrayLength(arrRef)
	stack.PushInt(arrLen)
}
