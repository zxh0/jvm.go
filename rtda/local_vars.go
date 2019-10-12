package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

type LocalVars struct {
	slots []heap.Slot
}

func newLocalVars(size uint) LocalVars {
	var slots []heap.Slot = nil
	if size > 0 {
		slots = make([]heap.Slot, size)
	}
	return LocalVars{slots: slots}
}

func (lv *LocalVars) GetIntVar(index uint) int32 {
	return lv.GetLocalVar(index).IntValue()
}
func (lv *LocalVars) SetIntVar(index uint, val int32) {
	lv.SetLocalVar(index, heap.NewIntSlot(val))
}

func (lv *LocalVars) GetLongVar(index uint) int64 {
	return lv.GetLocalVar(index).LongValue()
}
func (lv *LocalVars) SetLongVar(index uint, val int64) {
	lv.SetLocalVar(index, heap.NewLongSlot(val))
}

func (lv *LocalVars) GetFloatVar(index uint) float32 {
	return lv.GetLocalVar(index).FloatValue()
}
func (lv *LocalVars) SetFloatVar(index uint, val float32) {
	lv.SetLocalVar(index, heap.NewFloatSlot(val))
}

func (lv *LocalVars) GetDoubleVar(index uint) float64 {
	return lv.GetLocalVar(index).DoubleValue()
}
func (lv *LocalVars) SetDoubleVar(index uint, val float64) {
	lv.SetLocalVar(index, heap.NewDoubleSlot(val))
}

func (lv *LocalVars) GetRefVar(index uint) *heap.Object {
	return lv.GetLocalVar(index).Ref
}
func (lv *LocalVars) SetRefVar(index uint, ref *heap.Object) {
	lv.SetLocalVar(index, heap.NewRefSlot(ref))
}

func (lv *LocalVars) GetLocalVar(index uint) heap.Slot {
	return lv.slots[index]
}
func (lv *LocalVars) SetLocalVar(index uint, slot heap.Slot) {
	lv.slots[index] = slot
}

func (lv *LocalVars) GetBooleanVar(index uint) bool {
	return lv.GetIntVar(index) == 1
}
func (lv *LocalVars) GetThis() *heap.Object {
	return lv.GetRefVar(0)
}

func (lv *LocalVars) clearLocalVars() {
	for i := range lv.slots {
		lv.slots[i] = heap.EmptySlot
	}
}

func (lv *LocalVars) DebugGetSlots() []heap.Slot {
	return lv.slots
}
