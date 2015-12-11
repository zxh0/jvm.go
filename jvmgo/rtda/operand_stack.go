package rtda

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
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

func (self *OperandStack) IsEmpty() bool {
	return self.size == 0
}

func (self *OperandStack) PushNull() {
	self.slots[self.size] = nil
	self.size++
}

func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size] = ref
	self.size++
}
func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	top := self.slots[self.size]
	self.slots[self.size] = nil

	if top == nil {
		return nil
	} else {
		return top.(*heap.Object)
	}
}

func (self *OperandStack) PushBoolean(val bool) {
	if val {
		self.PushInt(1)
	} else {
		self.PushInt(0)
	}
}
func (self *OperandStack) PopBoolean() bool {
	return self.PopInt() == 1
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size] = val
	self.size++
}
func (self *OperandStack) PopInt() int32 {
	self.size--
	top := self.slots[self.size]
	self.slots[self.size] = nil
	return top.(int32)
}

// long consumes two slots
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size] = val
	self.size += 2
}
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	top := self.slots[self.size]
	self.slots[self.size] = nil
	return top.(int64)
}

func (self *OperandStack) PushFloat(val float32) {
	self.slots[self.size] = val
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	top := self.slots[self.size]
	self.slots[self.size] = nil
	return top.(float32)
}

// double consumes two slots
func (self *OperandStack) PushDouble(val float64) {
	self.slots[self.size] = val
	self.size += 2
}
func (self *OperandStack) PopDouble() float64 {
	self.size -= 2
	top := self.slots[self.size]
	self.slots[self.size] = nil
	return top.(float64)
}

func (self *OperandStack) PushSlot(any interface{}) {
	self.slots[self.size] = any
	self.size++
}
func (self *OperandStack) PopSlot() interface{} {
	self.size--
	top := self.slots[self.size]
	self.slots[self.size] = nil
	return top
}

func (self *OperandStack) PushField(any interface{}, isLongOrDouble bool) {
	self.slots[self.size] = any
	if isLongOrDouble {
		self.size += 2
	} else {
		self.size++
	}
}
func (self *OperandStack) PopField(isLongOrDouble bool) interface{} {
	if isLongOrDouble {
		self.size -= 2
	} else {
		self.size--
	}
	top := self.slots[self.size]
	self.slots[self.size] = nil
	return top
}

func (self *OperandStack) PopTops(n uint) []interface{} {
	start := self.size - n
	end := self.size
	top := self.slots[start:end]
	self.size -= n
	return top
}

func (self *OperandStack) TopRef(n uint) *heap.Object {
	ref := self.slots[self.size-1-n]
	if ref == nil {
		return nil
	} else {
		return ref.(*heap.Object)
	}
}

func (self *OperandStack) Clear() {
	self.size = 0
	for i := range self.slots {
		self.slots[i] = nil
	}
}

// only used by native methods
func (self *OperandStack) HackSetSlots(slots []interface{}) {
	self.slots = slots
	self.size = uint(len(slots))
}
