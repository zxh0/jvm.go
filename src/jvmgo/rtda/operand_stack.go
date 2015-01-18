package rtda

type OperandStack struct {
    size    int
    slots   []any
}

func (self *OperandStack) PushNull() {
    self.push(nil)
}

func (self *OperandStack) PushRef(ref *Ref) {
    self.push(ref)
}
func (self *OperandStack) PopRef() (*Ref) {
    return self.pop().(*Ref)
}

func (self *OperandStack) PushInt(val int32) {
    self.push(val)
}
func (self *OperandStack) PopInt() (int32) {
    return self.pop().(int32)
}

func (self *OperandStack) PushLong(val int64) {
    self.push(val)
}
func (self *OperandStack) PopLong() (int64) {
    return self.pop().(int64)
}

func (self *OperandStack) PushFloat(val float32) {
    self.push(val)
}
func (self *OperandStack) PopFloat() (float32) {
    return self.pop().(float32)
}

func (self *OperandStack) PushDouble(val float64) {
    self.push(val)
}
func (self *OperandStack) PopDouble() (float64) {
    return self.pop().(float64)
}

func (self *OperandStack) push(item any) {
    self.slots[self.size] = item
    self.size++
}
func (self *OperandStack) pop() (any) {
    self.size--
    return self.slots[self.size]
}

func newOperandStack(size uint16) (*OperandStack) {
    slots := make([]any, size)
    return &OperandStack{0, slots}
}
