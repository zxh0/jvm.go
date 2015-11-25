package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

//type Ref *rtda.Obj

// Load reference from array
type AALOAD struct{ base.NoOperandsInstruction }

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		ref := arrRef.Refs()[index]
		stack.PushRef(ref)
	}
}

// Load byte or boolean from array
type BALOAD struct{ base.NoOperandsInstruction }

func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Bytes()[index]
		stack.PushInt(int32(val))
	}
}

// Load char from array
type CALOAD struct{ base.NoOperandsInstruction }

func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Chars()[index]
		stack.PushInt(int32(val))
	}
}

// Load double from array
type DALOAD struct{ base.NoOperandsInstruction }

func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Doubles()[index]
		stack.PushDouble(val)
	}
}

// Load float from array
type FALOAD struct{ base.NoOperandsInstruction }

func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Floats()[index]
		stack.PushFloat(val)
	}
}

// Load int from array
type IALOAD struct{ base.NoOperandsInstruction }

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Ints()[index]
		stack.PushInt(val)
	}
}

// Load long from array
type LALOAD struct{ base.NoOperandsInstruction }

func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Longs()[index]
		stack.PushLong(val)
	}
}

// Load short from array
type SALOAD struct{ base.NoOperandsInstruction }

func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		val := arrRef.Shorts()[index]
		stack.PushInt(int32(val))
	}
}

func _aloadPop(frame *rtda.Frame) (*rtda.OperandStack, *rtc.Obj, int, bool) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if arrRef == nil {
		frame.Thread().ThrowNPE()
		return nil, nil, 0, false
	}
	if index < 0 || index >= rtc.ArrayLength(arrRef) {
		frame.Thread().ThrowArrayIndexOutOfBoundsException(index)
		return nil, nil, 0, false
	}

	return stack, arrRef, int(index), true
}
