package rtda

import (
    . "jvmgo/any"
    "jvmgo/rtda/class"
)

type OperandStack struct {
    size    uint
    slots   []Any
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

func (self *OperandStack) Push(item Any) {
    self.slots[self.size] = item
    self.size++
}
func (self *OperandStack) Pop() (Any) {
    self.size--
    return self.slots[self.size]
}

func (self *OperandStack) PopN(n uint) ([]Any) {
    start := self.size - n
    end := self.size - 1
    top := self.slots[start:end]
    self.size -= n
    return top
}

func newOperandStack(size uint16) (*OperandStack) {
    slots := make([]Any, size)
    return &OperandStack{0, slots}
}
