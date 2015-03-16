package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

//type Ref *rtda.Obj

// Load reference from array
type aaload struct{ NoOperandsInstruction }

func (self *aaload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		refArr := arrRef.Fields().([]*rtc.Obj)
		ref := refArr[index]
		stack.PushRef(ref)
	}
}

// Load byte or boolean from array
type baload struct{ NoOperandsInstruction }

func (self *baload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		byteArr := arrRef.Fields().([]int8)
		val := byteArr[index]
		stack.PushInt(int32(val))
	}
}

// Load char from array
type caload struct{ NoOperandsInstruction }

func (self *caload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		charArr := arrRef.Fields().([]uint16)
		val := charArr[index]
		stack.PushInt(int32(val))
	}
}

// Load double from array
type daload struct{ NoOperandsInstruction }

func (self *daload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		doubleArr := arrRef.Fields().([]float64)
		val := doubleArr[index]
		stack.PushDouble(val)
	}
}

// Load float from array
type faload struct{ NoOperandsInstruction }

func (self *faload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		floatArr := arrRef.Fields().([]float32)
		val := floatArr[index]
		stack.PushFloat(val)
	}
}

// Load int from array
type iaload struct{ NoOperandsInstruction }

func (self *iaload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		intArr := arrRef.Fields().([]int32)
		val := intArr[index]
		stack.PushInt(val)
	}
}

// Load long from array
type laload struct{ NoOperandsInstruction }

func (self *laload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		longArr := arrRef.Fields().([]int64)
		val := longArr[index]
		stack.PushLong(val)
	}
}

// Load short from array
type saload struct{ NoOperandsInstruction }

func (self *saload) Execute(frame *rtda.Frame) {
	stack, arrRef, index, ok := _aloadPop(frame)
	if ok {
		shortArr := arrRef.Fields().([]int16)
		val := shortArr[index]
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
