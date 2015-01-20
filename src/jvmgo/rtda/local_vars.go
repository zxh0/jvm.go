package rtda

import . "jvmgo/any"

type LocalVars struct {
    slots []Any
}

func (self *LocalVars) GetRef(index uint) (*Obj) {
    return self.slots[index].(*Obj)
}
func (self *LocalVars) SetRef(index uint, ref *Obj) {
    self.slots[index] = ref
}

func (self *LocalVars) GetInt(index uint) (int32) {
    return self.slots[index].(int32)
}
func (self *LocalVars) SetInt(index uint, val int32) {
    self.slots[index] = val
}

func (self *LocalVars) GetLong(index uint) (int64) {
    return self.slots[index].(int64)
}
func (self *LocalVars) SetLong(index uint, val int64) {
    self.slots[index] = val
}

func (self *LocalVars) GetFloat(index uint) (float32) {
    return self.slots[index].(float32)
}
func (self *LocalVars) SetFloat(index uint, val float32) {
    self.slots[index] = val
}

func (self *LocalVars) GetDouble(index uint) (float64) {
    return self.slots[index].(float64)
}
func (self *LocalVars) SetDouble(index uint, val float64) {
    self.slots[index] = val
}

func newLocalVars(size uint16) (*LocalVars) {
    slots := make([]Any, size)
    return &LocalVars{slots}
}
