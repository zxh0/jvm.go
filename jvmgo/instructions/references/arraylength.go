package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Get length of array
type arraylength struct{ base.NoOperandsInstruction }

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
