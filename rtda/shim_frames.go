package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

func newShimFrame(thread *Thread, args []Slot) *Frame {
	frame := &Frame{}
	frame.Thread = thread
	frame.Method = shimReturnMethod
	frame.OperandStack = OperandStack{
		size:  uint(len(args)),
		slots: args,
	}
	return frame
}

func newAthrowFrame(thread *Thread, ex *heap.Object, initArgs []Slot) *Frame {
	// stackSlots := [ex, ex, initArgs]
	stackSlots := make([]Slot, len(initArgs)+2)
	stackSlots[0] = heap.NewRefSlot(ex)
	stackSlots[1] = heap.NewRefSlot(ex)
	copy(stackSlots[2:], initArgs)

	frame := &Frame{}
	frame.Thread = thread
	frame.Method = shimAThrowMethod
	frame.OperandStack = OperandStack{
		size:  uint(len(stackSlots)),
		slots: stackSlots,
	}
	return frame
}
