package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// xstore: Store XXX into local variable
type Store struct {
	base.Index8Instruction
	L bool
}

func (instr *Store) Execute(frame *rtda.Frame) {
	frame.Store(instr.Index, instr.L)
}

// xstore_n: Store XXX into local variable
type StoreN struct {
	base.NoOperandsInstruction
	N uint
	L bool
}

func (instr *StoreN) Execute(frame *rtda.Frame) {
	frame.Store(instr.N, instr.L)
}

// Store into reference array
type AAStore struct{ base.NoOperandsInstruction }

func (instr *AAStore) Execute(frame *rtda.Frame) {
	ref := frame.PopRef()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Refs()[index] = ref
	}
}

// Store into byte or boolean array
type BAStore struct{ base.NoOperandsInstruction }

func (instr *BAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Bytes()[index] = int8(val)
	}
}

// Store into char array
type CAStore struct{ base.NoOperandsInstruction }

func (instr *CAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Chars()[index] = uint16(val)
	}
}

// Store into double array
type DAStore struct{ base.NoOperandsInstruction }

func (instr *DAStore) Execute(frame *rtda.Frame) {
	val := frame.PopDouble()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Doubles()[index] = val
	}
}

// Store into float array
type FAStore struct{ base.NoOperandsInstruction }

func (instr *FAStore) Execute(frame *rtda.Frame) {
	val := frame.PopFloat()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Floats()[index] = val
	}
}

// Store into int array
type IAStore struct{ base.NoOperandsInstruction }

func (instr *IAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Ints()[index] = val
	}
}

// Store into long array
type LAStore struct{ base.NoOperandsInstruction }

func (instr *LAStore) Execute(frame *rtda.Frame) {
	val := frame.PopLong()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Longs()[index] = val
	}
}

// Store into short array
type SAStore struct{ base.NoOperandsInstruction }

func (instr *SAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Shorts()[index] = int16(val)
	}
}

func _checkArrayAndIndex(frame *rtda.Frame, arrRef *heap.Object, index int32) bool {
	if arrRef == nil {
		frame.Thread.ThrowNPE()
		return false
	}
	if index < 0 || index >= heap.ArrayLength(arrRef) {
		frame.Thread.ThrowArrayIndexOutOfBoundsException(index)
		return false
	}
	return true
}
