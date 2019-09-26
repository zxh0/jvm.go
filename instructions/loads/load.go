package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// type Ref = *rtda.Object

// xload: Load XXX from local variable
type Load struct {
	base.Index8Instruction
	L bool
}

func (instr *Load) Execute(frame *rtda.Frame) {
	frame.Load(instr.Index, instr.L)
}

// xload_n: Load XXX from local variable
type LoadN struct {
	base.NoOperandsInstruction
	N uint
	L bool
}

func (instr *LoadN) Execute(frame *rtda.Frame) {
	frame.Load(instr.N, instr.L)
}

// Load reference from array
type AALoad struct{ base.NoOperandsInstruction }

func (instr *AALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		ref := arrRef.Refs()[index]
		stack.PushRef(ref)
	}
}

// Load byte or boolean from array
type BALoad struct{ base.NoOperandsInstruction }

func (instr *BALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Bytes()[index]
		stack.PushInt(int32(val))
	}
}

// Load char from array
type CALoad struct{ base.NoOperandsInstruction }

func (instr *CALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Chars()[index]
		stack.PushInt(int32(val))
	}
}

// Load double from array
type DALoad struct{ base.NoOperandsInstruction }

func (instr *DALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Doubles()[index]
		stack.PushDouble(val)
	}
}

// Load float from array
type FALoad struct{ base.NoOperandsInstruction }

func (instr *FALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Floats()[index]
		stack.PushFloat(val)
	}
}

// Load int from array
type IALoad struct{ base.NoOperandsInstruction }

func (instr *IALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Ints()[index]
		stack.PushInt(val)
	}
}

// Load long from array
type LALoad struct{ base.NoOperandsInstruction }

func (instr *LALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Longs()[index]
		stack.PushLong(val)
	}
}

// Load short from array
type SALoad struct{ base.NoOperandsInstruction }

func (instr *SALoad) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Shorts()[index]
		stack.PushInt(int32(val))
	}
}

func _aLoadPop(frame *rtda.Frame) (*rtda.OperandStack, *heap.Object, int, bool) {
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
