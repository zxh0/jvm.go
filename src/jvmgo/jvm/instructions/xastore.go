package instructions

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Store into reference array
type aastore struct{ NoOperandsInstruction }

func (self *aastore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	refArr := arrRef.Fields().([]*rtc.Obj)
	checkArrIndex(index, len(refArr))

	if val == nil {
		refArr[index] = nil
	} else {
		ref := val.(*rtc.Obj)
		refArr[index] = ref
	}
}

// Store into byte or boolean array
type bastore struct{ NoOperandsInstruction }

func (self *bastore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	byteArr := arrRef.Fields().([]int8)
	checkArrIndex(index, len(byteArr))
	byteArr[index] = int8(val.(int32))
}

// Store into char array
type castore struct{ NoOperandsInstruction }

func (self *castore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	charArr := arrRef.Fields().([]uint16)
	checkArrIndex(index, len(charArr))
	charArr[index] = uint16(val.(int32))
}

// Store into double array
type dastore struct{ NoOperandsInstruction }

func (self *dastore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	doubleArr := arrRef.Fields().([]float64)
	checkArrIndex(index, len(doubleArr))
	doubleArr[index] = val.(float64)
}

// Store into float array
type fastore struct{ NoOperandsInstruction }

func (self *fastore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	floatArr := arrRef.Fields().([]float32)
	checkArrIndex(index, len(floatArr))
	floatArr[index] = val.(float32)
}

// Store into int array
type iastore struct{ NoOperandsInstruction }

func (self *iastore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	intArr := arrRef.Fields().([]int32)
	checkArrIndex(index, len(intArr))
	intArr[index] = val.(int32)
}

// Store into long array
type lastore struct{ NoOperandsInstruction }

func (self *lastore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	longArr := arrRef.Fields().([]int64)
	checkArrIndex(index, len(longArr))
	longArr[index] = val.(int64)
}

// Store into short array
type sastore struct{ NoOperandsInstruction }

func (self *sastore) Execute(frame *rtda.Frame) {
	arrRef, index, val := popOperands(frame)
	shortArr := arrRef.Fields().([]int16)
	checkArrIndex(index, len(shortArr))
	shortArr[index] = int16(val.(int32))
}

func popOperands(frame *rtda.Frame) (*rtc.Obj, int, Any) {
	stack := frame.OperandStack()
	val := stack.Pop()
	index := int(stack.PopInt())
	arrRef := stack.PopRef()
	if arrRef == nil {
		// todo
		panic("NullPointerException")
	}
	return arrRef, index, val
}
