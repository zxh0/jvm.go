package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

type OperandStack struct {
	size  uint // TODO: change to int
	slots []Slot
}

func newOperandStack(size uint) *OperandStack {
	if size > 0 {
		slots := make([]Slot, size)
		return &OperandStack{0, slots}
	} else {
		return nil
	}
}

func (stack *OperandStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *OperandStack) PushNull() {
	stack.PushSlot(EmptySlot)
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
	stack.PushSlot(heap.NewIntSlot(val))
}
func (stack *OperandStack) PopInt() int32 {
	return stack.PopSlot().IntValue()
}

// long consumes two slots
func (stack *OperandStack) PushLong(val int64) {
	stack.PushSlot(heap.NewLongSlot(val))
	stack.size++
}
func (stack *OperandStack) PopLong() int64 {
	stack.size--
	return stack.PopSlot().LongValue()
}

func (stack *OperandStack) PushFloat(val float32) {
	stack.PushSlot(heap.NewFloatSlot(val))
}
func (stack *OperandStack) PopFloat() float32 {
	return stack.PopSlot().FloatValue()
}

// double consumes two slots
func (stack *OperandStack) PushDouble(val float64) {
	stack.PushSlot(heap.NewDoubleSlot(val))
	stack.size++
}
func (stack *OperandStack) PopDouble() float64 {
	stack.size--
	return stack.PopSlot().DoubleValue()
}

func (stack *OperandStack) PushRef(ref *heap.Object) {
	stack.PushSlot(heap.NewRefSlot(ref))
}
func (stack *OperandStack) PopRef() *heap.Object {
	return stack.PopSlot().Ref
}

func (stack *OperandStack) PushSlot(slot Slot) {
	stack.slots[stack.size] = slot
	stack.size++
}
func (stack *OperandStack) PopSlot() Slot {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = EmptySlot // help GC
	return top
}

func (stack *OperandStack) PushField(slot Slot, isLongOrDouble bool) {
	stack.PushSlot(slot)
	if isLongOrDouble {
		stack.size++
	}
}
func (stack *OperandStack) PopField(isLongOrDouble bool) Slot {
	if isLongOrDouble {
		stack.size--
	}
	return stack.PopSlot()
}

func (stack *OperandStack) PopTops(n uint) []Slot {
	start := stack.size - n
	end := stack.size
	top := stack.slots[start:end]
	stack.size -= n
	return top
}

func (stack *OperandStack) TopRef(n uint) *heap.Object {
	return stack.slots[stack.size-1-n].Ref
}

func (stack *OperandStack) Clear() {
	stack.size = 0
	for i := range stack.slots {
		stack.slots[i] = EmptySlot
	}
}

// only used by native methods
func (stack *OperandStack) HackSetSlots(slots []Slot) {
	stack.slots = slots
	stack.size = uint(len(slots))
}
