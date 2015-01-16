package rtda

type OperandStack struct {
    size    uint16
    slots   []any
}

func (self *OperandStack) push(item any) {
    self.slots[self.size] = item
    self.size++
}
func (self *OperandStack) pop() (item any) {
    item = self.slots[self.size]
    self.size--
    return
} 

func (self *OperandStack) pushInt(val int32) {
    self.push(val)
}
func (self *OperandStack) popInt() (int32) {
    return self.pop().(int32)
}

func (self *OperandStack) pushFloat(val float32) {
    self.push(val)
}
func (self *OperandStack) popFloat() (float32) {
    return self.pop().(float32)
} 



func NewOperandStack(length uint16) (*OperandStack) {
    slots := make([]any, length)
    return &OperandStack{0, slots}
}
