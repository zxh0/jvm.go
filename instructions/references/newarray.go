package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Create new array
type NewArray struct {
	atype uint8
}

func (instr *NewArray) FetchOperands(reader *base.CodeReader) {
	instr.atype = reader.ReadUint8()
}
func (instr *NewArray) Execute(frame *rtda.Frame) {
	count := frame.PopInt()
	if count < 0 {
		frame.Thread.ThrowNegativeArraySizeException()
		return
	}

	arr := frame.GetRuntime().NewPrimitiveArray(instr.atype, uint(count))
	frame.PushRef(arr)
}

// Create new array of reference
type ANewArray struct{ base.Index16Instruction }

func (instr *ANewArray) Execute(frame *rtda.Frame) {
	cp := frame.GetConstantPool()
	kClass := cp.GetConstantClass(instr.Index)
	componentClass := kClass.GetClass()

	if componentClass.InitializationNotStarted() {
		thread := frame.Thread
		frame.NextPC = thread.PC // undo anewarray
		thread.InitClass(componentClass)
		return
	}

	count := frame.PopInt()
	if count < 0 {
		frame.Thread.ThrowNegativeArraySizeException()
	} else {
		arr := componentClass.NewArray(uint(count))
		frame.PushRef(arr)
	}
}

// Create new multidimensional array
type MultiANewArray struct {
	index      uint16
	dimensions uint8
}

func (instr *MultiANewArray) FetchOperands(reader *base.CodeReader) {
	instr.index = reader.ReadUint16()
	instr.dimensions = reader.ReadUint8()
}
func (instr *MultiANewArray) Execute(frame *rtda.Frame) {
	cp := frame.GetConstantPool()
	kClass := cp.GetConstantClass(uint(instr.index))
	arrClass := kClass.GetClass()

	counts := frame.PopTops(uint(instr.dimensions))
	if !_checkCounts(counts) {
		frame.Thread.ThrowNegativeArraySizeException()
	} else {
		arr := _newMultiArray(counts, arrClass)
		frame.PushRef(arr)
	}
}

func _checkCounts(counts []heap.Slot) bool {
	for _, c := range counts {
		if c.IntValue() < 0 {
			return false
		}
	}
	return true
}

func _newMultiArray(counts []heap.Slot, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0].IntValue())
	arr := heap.NewArray(arrClass, count)

	if len(counts) > 1 {
		objs := arr.GetRefs()
		for i := range objs {
			objs[i] = _newMultiArray(counts[1:], arrClass.GetComponentClass())
		}
	}

	return arr
}
