package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

type LocalVars struct {
	slots []Slot
}

func newLocalVars(size uint) *LocalVars {
	if size > 0 {
		slots := make([]Slot, size)
		return &LocalVars{slots}
	} else {
		return nil
	}
}

func (lv *LocalVars) GetInt(index uint) int32 {
	return lv.Get(index).IntValue()
}
func (lv *LocalVars) SetInt(index uint, val int32) {
	lv.Set(index, heap.NewIntSlot(val))
}

func (lv *LocalVars) GetLong(index uint) int64 {
	return lv.Get(index).LongValue()
}
func (lv *LocalVars) SetLong(index uint, val int64) {
	lv.Set(index, heap.NewLongSlot(val))
}

func (lv *LocalVars) GetFloat(index uint) float32 {
	return lv.Get(index).FloatValue()
}
func (lv *LocalVars) SetFloat(index uint, val float32) {
	lv.Set(index, heap.NewFloatSlot(val))
}

func (lv *LocalVars) GetDouble(index uint) float64 {
	return lv.Get(index).DoubleValue()
}
func (lv *LocalVars) SetDouble(index uint, val float64) {
	lv.Set(index, heap.NewDoubleSlot(val))
}

func (lv *LocalVars) GetRef(index uint) *heap.Object {
	return lv.Get(index).Ref
}
func (lv *LocalVars) SetRef(index uint, ref *heap.Object) {
	lv.Set(index, heap.NewRefSlot(ref))
}

func (lv *LocalVars) Get(index uint) Slot {
	return lv.slots[index]
}
func (lv *LocalVars) Set(index uint, slot Slot) {
	lv.slots[index] = slot
}

func (lv *LocalVars) GetBoolean(index uint) bool {
	return lv.GetInt(index) == 1
}
func (lv *LocalVars) GetThis() *heap.Object {
	return lv.GetRef(0)
}

func (lv *LocalVars) clear() {
	for i := range lv.slots {
		lv.slots[i] = EmptySlot
	}
}
