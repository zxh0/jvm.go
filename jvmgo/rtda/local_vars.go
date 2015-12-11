package rtda

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

type LocalVars struct {
	slots []interface{}
}

func newLocalVars(size uint) *LocalVars {
	if size > 0 {
		slots := make([]interface{}, size)
		return &LocalVars{slots}
	} else {
		return nil
	}
}

func (self *LocalVars) GetThis() *heap.Object {
	return self.GetRef(0)
}

func (self *LocalVars) GetRef(index uint) *heap.Object {
	ref := self.slots[index]
	if ref == nil {
		return nil
	} else {
		return ref.(*heap.Object)
	}
}
func (self *LocalVars) SetRef(index uint, ref *heap.Object) {
	self.slots[index] = ref
}

func (self *LocalVars) GetBoolean(index uint) bool {
	return self.GetInt(index) == 1
}

func (self *LocalVars) GetInt(index uint) int32 {
	return self.slots[index].(int32)
}
func (self *LocalVars) SetInt(index uint, val int32) {
	self.slots[index] = val
}

func (self *LocalVars) GetLong(index uint) int64 {
	return self.slots[index].(int64)
}
func (self *LocalVars) SetLong(index uint, val int64) {
	self.slots[index] = val
}

func (self *LocalVars) GetFloat(index uint) float32 {
	return self.slots[index].(float32)
}
func (self *LocalVars) SetFloat(index uint, val float32) {
	self.slots[index] = val
}

func (self *LocalVars) GetDouble(index uint) float64 {
	return self.slots[index].(float64)
}
func (self *LocalVars) SetDouble(index uint, val float64) {
	self.slots[index] = val
}

func (self *LocalVars) Get(index uint) interface{} {
	return self.slots[index]
}
func (self *LocalVars) Set(index uint, any interface{}) {
	self.slots[index] = any
}

func (self *LocalVars) clear() {
	for i := range self.slots {
		self.slots[i] = nil
	}
}
