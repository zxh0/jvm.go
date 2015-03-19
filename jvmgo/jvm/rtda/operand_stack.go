package rtda

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

type OperandStack struct {
	size  uint
	slots []Any
}

func newOperandStack(size uint) *OperandStack {
	if size > 0 {
		slots := make([]Any, size)
		return &OperandStack{0, slots}
	} else {
		return nil
	}
}

func (self *OperandStack) IsEmpty() bool {
	return self.size == 0
}

func (self *OperandStack) PushNull() {
	self.Push(nil)
}

func (self *OperandStack) PushRef(ref *rtc.Obj) {
	self.Push(ref)
}
func (self *OperandStack) PopRef() *rtc.Obj {
	ref := self.Pop()
	if ref == nil {
		return nil
	} else {
		return ref.(*rtc.Obj)
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
	self.Push(val)
}
func (self *OperandStack) PopInt() int32 {
	return self.Pop().(int32)
}

func (self *OperandStack) PushLong(val int64) {
	self.Push(val)
}
func (self *OperandStack) PopLong() int64 {
	return self.Pop().(int64)
}

func (self *OperandStack) PushFloat(val float32) {
	self.Push(val)
}
func (self *OperandStack) PopFloat() float32 {
	return self.Pop().(float32)
}

func (self *OperandStack) PushDouble(val float64) {
	self.Push(val)
}
func (self *OperandStack) PopDouble() float64 {
	return self.Pop().(float64)
}

func (self *OperandStack) Push(any Any) {
	self.slots[self.size] = any
	self.size++
}
func (self *OperandStack) Pop() Any {
	self.size--
	top := self.slots[self.size]
	self.slots[self.size] = nil
	return top
}

func (self *OperandStack) PopTops(n uint) []Any {
	start := self.size - n
	end := self.size
	top := self.slots[start:end]
	self.size -= n
	return top
}

func (self *OperandStack) Top(n uint) Any {
	return self.slots[self.size-1-n]
}
func (self *OperandStack) TopRef(n uint) *rtc.Obj {
	ref := self.slots[self.size-1-n]
	if ref == nil {
		return nil
	} else {
		return ref.(*rtc.Obj)
	}
}

func (self *OperandStack) Clear() {
	for !self.IsEmpty() {
		self.Pop()
	}
}

// only used by native methods
func (self *OperandStack) HackSetSlots(slots []Any) {
	self.slots = slots
	self.size = uint(len(slots))
}
