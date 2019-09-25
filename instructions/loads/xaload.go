package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

//type Ref *rtda.Object

// Load reference from array
type AALOAD struct{ base.NoOperandsInstruction }

func (instr *AALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		ref := arrRef.Refs()[index]
		stack.PushRef(ref)
	}
}

// Load byte or boolean from array
type BALOAD struct{ base.NoOperandsInstruction }

func (instr *BALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Bytes()[index]
		stack.PushInt(int32(val))
	}
}

// Load char from array
type CALOAD struct{ base.NoOperandsInstruction }

func (instr *CALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Chars()[index]
		stack.PushInt(int32(val))
	}
}

// Load double from array
type DALOAD struct{ base.NoOperandsInstruction }

func (instr *DALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Doubles()[index]
		stack.PushDouble(val)
	}
}

// Load float from array
type FALOAD struct{ base.NoOperandsInstruction }

func (instr *FALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Floats()[index]
		stack.PushFloat(val)
	}
}

// Load int from array
type IALOAD struct{ base.NoOperandsInstruction }

func (instr *IALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Ints()[index]
		stack.PushInt(val)
	}
}

// Load long from array
type LALOAD struct{ base.NoOperandsInstruction }

func (instr *LALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Longs()[index]
		stack.PushLong(val)
	}
}

// Load short from array
type SALOAD struct{ base.NoOperandsInstruction }

func (instr *SALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Shorts()[index]
		stack.PushInt(int32(val))
	}
}

func _aloadPop(frame *rtda.Frame) (*rtda.OperandStack, *heap.Object, int, bool) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if arrRef == nil {
		frame.Thread().ThrowNPE()
		return nil, nil, 0, false
	}
	if index < 0 || index >= heap.ArrayLength(arrRef) {
		frame.Thread().ThrowArrayIndexOutOfBoundsException(index)
		return nil, nil, 0, false
	}

	return stack, arrRef, int(index), true
}
