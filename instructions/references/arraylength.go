package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Get length of array
type ArrayLength struct{ base.NoOperandsInstruction }

func (instr *ArrayLength) Execute(frame *rtda.Frame) {
	arrRef := frame.PopRef()

	if arrRef == nil {
		frame.Thread().ThrowNPE()
		return
	}

	arrLen := heap.ArrayLength(arrRef)
	frame.PushInt(arrLen)
}
