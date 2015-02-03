package rtda

import (
    . "jvmgo/any"
    "jvmgo/rtda/class"
)

type OperandStack struct {
    size    uint
    slots   []Any
}

func newOperandStack(size uint) (*OperandStack) {
    slots := make([]Any, size)
    return &OperandStack{0, slots}
}

func (self *OperandStack) PushNull() {
    self.Push(nil)
}

func (self *OperandStack) PushRef(ref *class.Obj) {
    self.Push(ref)
}
func (self *OperandStack) PopRef() (*class.Obj) {
    ref := self.Pop()
    if ref == nil {
        return nil
    } else {
        return ref.(*class.Obj)
    }
}

func (self *OperandStack) PushBoolean(val bool) {
    if val {
        self.PushInt(1)
    } else {
        self.PushInt(0)
    }
}
func (self *OperandStack) PopBoolean() (bool) {
    return self.PopInt() == 1
}

func (self *OperandStack) PushInt(val int32) {
    self.Push(val)
}
func (self *OperandStack) PopInt() (int32) {
    return self.Pop().(int32)
}

func (self *OperandStack) PushLong(val int64) {
    self.Push(val)
}
func (self *OperandStack) PopLong() (int64) {
    return self.Pop().(int64)
}

func (self *OperandStack) PushFloat(val float32) {
    self.Push(val)
}
func (self *OperandStack) PopFloat() (float32) {
    return self.Pop().(float32)
}

func (self *OperandStack) PushDouble(val float64) {
    self.Push(val)
}
func (self *OperandStack) PopDouble() (float64) {
    return self.Pop().(float64)
}

func (self *OperandStack) Push(any Any) {
    self.slots[self.size] = any
    self.size++
}
func (self *OperandStack) Pop() (Any) {
    self.size--
    top := self.slots[self.size]
    self.slots[self.size] = nil
    return top
}

func (self *OperandStack) PopN(n uint) ([]Any) {
    start := self.size - n
    end := self.size
    top := self.slots[start:end]
    self.size -= n
    return top
}
func (self *OperandStack) Top(n uint) (Any) {
    return self.slots[self.size - 1 - n]
}
