package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

type OperandStack struct {
	size  uint
	slots []interface{}
}

func newOperandStack(size uint) *OperandStack {
	if size > 0 {
		slots := make([]interface{}, size)
		return &OperandStack{0, slots}
	} else {
		return nil
	}
}

func (stack *OperandStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *OperandStack) PushNull() {
	stack.slots[stack.size] = nil
	stack.size++
}

func (stack *OperandStack) PushRef(ref *heap.Object) {
	stack.slots[stack.size] = ref
	stack.size++
}
func (stack *OperandStack) PopRef() *heap.Object {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = nil

	if top == nil {
		return nil
	} else {
		return top.(*heap.Object)
	}
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
	stack.slots[stack.size] = val
	stack.size++
}
func (stack *OperandStack) PopInt() int32 {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = nil
	return top.(int32)
}

// long consumes two slots
func (stack *OperandStack) PushLong(val int64) {
	stack.slots[stack.size] = val
	stack.size += 2
}
func (stack *OperandStack) PopLong() int64 {
	stack.size -= 2
	top := stack.slots[stack.size]
	stack.slots[stack.size] = nil
	return top.(int64)
}

func (stack *OperandStack) PushFloat(val float32) {
	stack.slots[stack.size] = val
	stack.size++
}
func (stack *OperandStack) PopFloat() float32 {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = nil
	return top.(float32)
}

// double consumes two slots
func (stack *OperandStack) PushDouble(val float64) {
	stack.slots[stack.size] = val
	stack.size += 2
}
func (stack *OperandStack) PopDouble() float64 {
	stack.size -= 2
	top := stack.slots[stack.size]
	stack.slots[stack.size] = nil
	return top.(float64)
}

func (stack *OperandStack) PushSlot(any interface{}) {
	stack.slots[stack.size] = any
	stack.size++
}
func (stack *OperandStack) PopSlot() interface{} {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = nil
	return top
}

func (stack *OperandStack) PushField(any interface{}, isLongOrDouble bool) {
	stack.slots[stack.size] = any
	if isLongOrDouble {
		stack.size += 2
	} else {
		stack.size++
	}
}
func (stack *OperandStack) PopField(isLongOrDouble bool) interface{} {
	if isLongOrDouble {
		stack.size -= 2
	} else {
		stack.size--
	}
	top := stack.slots[stack.size]
	stack.slots[stack.size] = nil
	return top
}

func (stack *OperandStack) PopTops(n uint) []interface{} {
	start := stack.size - n
	end := stack.size
	top := stack.slots[start:end]
	stack.size -= n
	return top
}

func (stack *OperandStack) TopRef(n uint) *heap.Object {
	ref := stack.slots[stack.size-1-n]
	if ref == nil {
		return nil
	} else {
		return ref.(*heap.Object)
	}
}

func (stack *OperandStack) Clear() {
	stack.size = 0
	for i := range stack.slots {
		stack.slots[i] = nil
	}
}

// only used by native methods
func (stack *OperandStack) HackSetSlots(slots []interface{}) {
	stack.slots = slots
	stack.size = uint(len(slots))
}
