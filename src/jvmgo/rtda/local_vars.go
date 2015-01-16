package rtda

type LocalVars struct {
    slots []any
}

func (self *LocalVars) set(index uint16, item any) {
    self.slots[index] = item
}

func (self *LocalVars) getInt(index uint16) (int32) {
    return self.slots[index].(int32)
}

func (self *LocalVars) getLong(index uint16) (int64) {
    return self.slots[index].(int64)
}

func (self *LocalVars) getFloat(index uint16) (float32) {
    return self.slots[index].(float32)
}

func (self *LocalVars) getDouble(index uint16) (float64) {
    return self.slots[index].(float64)
}

// func (self *LocalVars) getRef(index uint16)

func newLocalVars(length uint16) (*LocalVars) {
    slots := make([]any, length)
    return &LocalVars{slots}
}
