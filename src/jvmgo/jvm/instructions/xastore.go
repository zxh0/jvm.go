package instructions

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Store into reference array
type aastore struct{ NoOperandsInstruction }

func (self *aastore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		refArr := arrRef.Fields().([]*rtc.Obj)
		if val == nil {
			refArr[index] = nil
		} else {
			ref := val.(*rtc.Obj)
			refArr[index] = ref
		}
	}
}

// Store into byte or boolean array
type bastore struct{ NoOperandsInstruction }

func (self *bastore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		byteArr := arrRef.Fields().([]int8)
		byteArr[index] = int8(val.(int32))
	}
}

// Store into char array
type castore struct{ NoOperandsInstruction }

func (self *castore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		charArr := arrRef.Fields().([]uint16)
		charArr[index] = uint16(val.(int32))
	}
}

// Store into double array
type dastore struct{ NoOperandsInstruction }

func (self *dastore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		doubleArr := arrRef.Fields().([]float64)
		doubleArr[index] = val.(float64)
	}
}

// Store into float array
type fastore struct{ NoOperandsInstruction }

func (self *fastore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		floatArr := arrRef.Fields().([]float32)
		floatArr[index] = val.(float32)
	}
}

// Store into int array
type iastore struct{ NoOperandsInstruction }

func (self *iastore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		intArr := arrRef.Fields().([]int32)
		intArr[index] = val.(int32)
	}
}

// Store into long array
type lastore struct{ NoOperandsInstruction }

func (self *lastore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		longArr := arrRef.Fields().([]int64)
		longArr[index] = val.(int64)
	}
}

// Store into short array
type sastore struct{ NoOperandsInstruction }

func (self *sastore) Execute(frame *rtda.Frame) {
	arrRef, index, val, ok := _astorePop(frame)
	if ok {
		shortArr := arrRef.Fields().([]int16)
		shortArr[index] = int16(val.(int32))
	}
}

func _astorePop(frame *rtda.Frame) (*rtc.Obj, int, Any, bool) {
	stack := frame.OperandStack()
	val := stack.Pop()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if arrRef == nil {
		frame.Thread().ThrowNPE()
		return nil, 0, nil, false
	}
	if index < 0 || index >= rtc.ArrayLength(arrRef) {
		frame.Thread().ThrowArrayIndexOutOfBoundsException(index)
		return nil, 0, nil, false
	}

	return arrRef, int(index), val, true
}
