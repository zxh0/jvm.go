package rtda

type OperandStack struct {
    size    int
    slots   []any
}

func (self *OperandStack) PushNull() {
    self.Push(nil)
}

func (self *OperandStack) PushRef(ref *Ref) {
    self.Push(ref)
}
func (self *OperandStack) PopRef() (*Ref) {
    return self.Pop().(*Ref)
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

func (self *OperandStack) Push(item any) {
    self.slots[self.size] = item
    self.size++
}
func (self *OperandStack) Pop() (any) {
    self.size--
    return self.slots[self.size]
}

func newOperandStack(size uint16) (*OperandStack) {
    slots := make([]any, size)
    return &OperandStack{0, slots}
}
