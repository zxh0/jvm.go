package rtda

type OperandStack struct {
    size    uint16
    slots   []any
}

func (self *OperandStack) pushInt(val int32) {
    self.slots[self.size] = val
    self.size++
}
func (self *OperandStack) popInt() (val int32) {
    val = self.slots[self.size].(int32)
    self.size--
    return
}

func NewOperandStack(length uint16) (*OperandStack) {
    slots := make([]any, length)
    return &OperandStack{0, slots}
}
