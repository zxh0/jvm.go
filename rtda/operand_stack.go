package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

type OperandStack struct {
	size  uint // TODO: change to int
	slots []heap.Slot
}

func newOperandStackWithSlots(slots []heap.Slot) OperandStack {
	return OperandStack{
		size:  uint(len(slots)),
		slots: slots,
	}
}

func newOperandStack(size uint) OperandStack {
	var slots []heap.Slot = nil
	if size > 0 {
		slots = make([]heap.Slot, size)
	}
	return OperandStack{size: 0, slots: slots}
}

func (stack *OperandStack) IsStackEmpty() bool {
	return stack.size == 0
}

func (stack *OperandStack) PushNull() {
	stack.Push(heap.EmptySlot)
}

func (stack *OperandStack) PushBoolean(val bool) {
	if val {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
func (stack *OperandStack) PopBoolean() bool {
	return stack.PopInt() == 1
}

func (stack *OperandStack) PushInt(val int32) {
	stack.Push(heap.NewIntSlot(val))
}
func (stack *OperandStack) PopInt() int32 {
	return stack.Pop().IntValue()
}

// long consumes two slots
func (stack *OperandStack) PushLong(val int64) {
	stack.Push(heap.NewLongSlot(val))
	stack.size++
}
func (stack *OperandStack) PopLong() int64 {
	stack.size--
	return stack.Pop().LongValue()
}

func (stack *OperandStack) PushFloat(val float32) {
	stack.Push(heap.NewFloatSlot(val))
}
func (stack *OperandStack) PopFloat() float32 {
	return stack.Pop().FloatValue()
}

// double consumes two slots
func (stack *OperandStack) PushDouble(val float64) {
	stack.Push(heap.NewDoubleSlot(val))
	stack.size++
}
func (stack *OperandStack) PopDouble() float64 {
	stack.size--
	return stack.Pop().DoubleValue()
}

func (stack *OperandStack) PushRef(ref *heap.Object) {
	stack.Push(heap.NewRefSlot(ref))
}
func (stack *OperandStack) PopRef() *heap.Object {
	return stack.Pop().Ref
}

func (stack *OperandStack) Push(slot heap.Slot) {
	stack.slots[stack.size] = slot
	stack.size++
}
func (stack *OperandStack) Pop() heap.Slot {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = heap.EmptySlot // help GC
	return top
}

func (stack *OperandStack) PushL(slot heap.Slot, isLongOrDouble bool) {
	stack.Push(slot)
	if isLongOrDouble {
		stack.size++
	}
}
func (stack *OperandStack) PopL(isLongOrDouble bool) heap.Slot {
	if isLongOrDouble {
		stack.size--
	}
	return stack.Pop()
}

func (stack *OperandStack) PopTops(n uint) []heap.Slot {
	start := stack.size - n
	end := stack.size
	top := stack.slots[start:end]
	stack.size -= n
	return top
}

func (stack *OperandStack) TopRef(n uint) *heap.Object {
	return stack.slots[stack.size-1-n].Ref
}

func (stack *OperandStack) ClearStack() {
	stack.size = 0
	for i := range stack.slots {
		stack.slots[i] = heap.EmptySlot
	}
}

// only used by native methods
func (stack *OperandStack) HackSetSlots(slots []heap.Slot) { // TODO
	stack.slots = slots
	stack.size = uint(len(slots))
}

func (stack *OperandStack) DebugGetSlots() []heap.Slot {
	return stack.slots[:stack.size]
}
