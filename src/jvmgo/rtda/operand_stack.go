package rtda

type OperandStack struct {
    size    uint16
    slots   []any
}

func (self *OperandStack) pushInt(val int32) {
    self.slots[self.size] = val
    self.size++
}

func (self *OperandStack) popInt() (int32) {
    return self.slots[self.size].(int32)
}

func NewOperandStack(length uint16) (*OperandStack) {
    slots := make([]any, length)
    return &OperandStack{0, slots}
}
