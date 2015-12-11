package rtda

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func newShimFrame(thread *Thread, args []interface{}) *Frame {
	frame := &Frame{}
	frame.thread = thread
	frame.method = heap.ReturnMethod()
	frame.operandStack = &OperandStack{
		size:  uint(len(args)),
		slots: args,
	}
	return frame
}

func newAthrowFrame(thread *Thread, ex *heap.Object, initArgs []interface{}) *Frame {
	// stackSlots := [ex, ex, initArgs]
	stackSlots := make([]interface{}, len(initArgs)+2)
	stackSlots[0] = ex
	stackSlots[1] = ex
	copy(stackSlots[2:], initArgs)

	frame := &Frame{}
	frame.thread = thread
	frame.method = heap.AthrowMethod()
	frame.operandStack = &OperandStack{
		size:  uint(len(stackSlots)),
		slots: stackSlots,
	}
	return frame
}
