package rtda

// todo
type any interface{}

// used by LocalVars and OperandStack
type Slots struct {
    slots []any
}

func (self *Slots) push(item any) {
    // todo
}

func (self *Slots) pop() (any) {
    // todo
    return nil
}
