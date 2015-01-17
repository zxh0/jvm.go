package rtda

type OperandStack struct {
    size    int
    slots   []any
}

func (self *OperandStack) push(item any) {
    self.slots[self.size] = item
    self.size++
}

func (self *OperandStack) pop() (any) {
    self.size--
    return self.slots[self.size]
}

func (self *OperandStack) PushNull() {
    self.push(nil)
}

func (self *OperandStack) popInt() (int32) {
    return self.pop().(int32)
}

func (self *OperandStack) popLong() (int64) {
    return self.pop().(int64)
}

func (self *OperandStack) popFloat() (float32) {
    return self.pop().(float32)
}

func (self *OperandStack) popDouble() (float64) {
    return self.pop().(float64)
}

func (self *OperandStack) popRef() (*Ref) {
    return self.pop().(*Ref)
}

func newOperandStack(size uint16) (*OperandStack) {
    slots := make([]any, size)
    return &OperandStack{0, slots}
}
