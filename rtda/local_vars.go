package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
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

func (lv *LocalVars) GetThis() *heap.Object {
	return lv.GetRef(0)
}

func (lv *LocalVars) GetRef(index uint) *heap.Object {
	ref := lv.slots[index]
	if ref == nil {
		return nil
	} else {
		return ref.(*heap.Object)
	}
}
func (lv *LocalVars) SetRef(index uint, ref *heap.Object) {
	lv.slots[index] = ref
}

func (lv *LocalVars) GetBoolean(index uint) bool {
	return lv.GetInt(index) == 1
}

func (lv *LocalVars) GetInt(index uint) int32 {
	return lv.slots[index].(int32)
}
func (lv *LocalVars) SetInt(index uint, val int32) {
	lv.slots[index] = val
}

func (lv *LocalVars) GetLong(index uint) int64 {
	return lv.slots[index].(int64)
}
func (lv *LocalVars) SetLong(index uint, val int64) {
	lv.slots[index] = val
}

func (lv *LocalVars) GetFloat(index uint) float32 {
	return lv.slots[index].(float32)
}
func (lv *LocalVars) SetFloat(index uint, val float32) {
	lv.slots[index] = val
}

func (lv *LocalVars) GetDouble(index uint) float64 {
	return lv.slots[index].(float64)
}
func (lv *LocalVars) SetDouble(index uint, val float64) {
	lv.slots[index] = val
}

func (lv *LocalVars) Get(index uint) interface{} {
	return lv.slots[index]
}
func (lv *LocalVars) Set(index uint, any interface{}) {
	lv.slots[index] = any
}

func (lv *LocalVars) clear() {
	for i := range lv.slots {
		lv.slots[i] = nil
	}
}
