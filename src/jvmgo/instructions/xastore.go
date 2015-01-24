package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Store into reference array 
type aastore struct {NoOperandsInstruction}
func (self *aastore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    refArr := arrRef.Fields().([]*class.Obj)
    checkArrIndex(index, len(refArr))
    // todo
    ref := val.(*class.Obj)
    refArr[index] = ref
}

// Store into byte or boolean array 
type bastore struct {NoOperandsInstruction}
func (self *bastore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    byteArr := arrRef.Fields().([]int8)
    checkArrIndex(index, len(byteArr))
    byteArr[index] = val.(int8)
}

// Store into char array 
type castore struct {NoOperandsInstruction}
func (self *castore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    charArr := arrRef.Fields().([]uint16)
    checkArrIndex(index, len(charArr))
    charArr[index] = val.(uint16)
}

// Store into double array 
type dastore struct {NoOperandsInstruction}
func (self *dastore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    doubleArr := arrRef.Fields().([]float64)
    checkArrIndex(index, len(doubleArr))
    doubleArr[index] = val.(float64)
}

// Store into float array 
type fastore struct {NoOperandsInstruction}
func (self *fastore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    floatArr := arrRef.Fields().([]float32)
    checkArrIndex(index, len(floatArr))
    floatArr[index] = val.(float32)
}

// Store into int array 
type iastore struct {NoOperandsInstruction}
func (self *iastore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    intArr := arrRef.Fields().([]int32)
    checkArrIndex(index, len(intArr))
    intArr[index] = val.(int32)
}

// Store into long array 
type lastore struct {NoOperandsInstruction}
func (self *lastore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    longArr := arrRef.Fields().([]int64)
    checkArrIndex(index, len(longArr))
    longArr[index] = val.(int64)
}

// Store into short array 
type sastore struct {NoOperandsInstruction}
func (self *sastore) Execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    shortArr := arrRef.Fields().([]int16)
    checkArrIndex(index, len(shortArr))
    shortArr[index] = val.(int16)
}

func popOperands(thread *rtda.Thread) (*class.Obj, int, Any) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.Pop()
    index := int(stack.PopInt())
    arrRef := stack.PopRef()
    if arrRef == nil {
        // todo
        panic("NullPointerException")
    }
    return arrRef, index, val
}
