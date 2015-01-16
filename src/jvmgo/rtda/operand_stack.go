package rtda

type OperandStack struct {
    slots []any
}

func (self *OperandStack) push(item any) {
    // todo
}

func (self *OperandStack) pop() (any) {
    // todo
    return nil
}

func NewOperandStack(length uint16) (*OperandStack) {
    slots := make([]any, length)
    return &OperandStack{slots}
}
