package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

//type Ref *rtda.Obj

// Load reference from array
type aaload struct {NoOperandsInstruction}
func (self *aaload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    refArr := arrRef.Fields().([]*class.Obj)
    checkArrIndex(index, len(refArr))
    ref := refArr[index]
    stack.PushRef(ref)
}

// Load byte or boolean from array 
type baload struct {NoOperandsInstruction}
func (self *baload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    byteArr := arrRef.Fields().([]int8)
    checkArrIndex(index, len(byteArr))
    val := byteArr[index]
    stack.PushInt(int32(val))
}

// Load char from array 
type caload struct {NoOperandsInstruction}
func (self *caload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    charArr := arrRef.Fields().([]uint16)
    checkArrIndex(index, len(charArr))
    val := charArr[index]
    stack.PushInt(int32(val))
}

// Load double from array 
type daload struct {NoOperandsInstruction}
func (self *daload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    doubleArr := arrRef.Fields().([]float64)
    checkArrIndex(index, len(doubleArr))
    val := doubleArr[index]
    stack.PushDouble(val)
}

// Load float from array 
type faload struct {NoOperandsInstruction}
func (self *faload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    floatArr := arrRef.Fields().([]float32)
    checkArrIndex(index, len(floatArr))
    val := floatArr[index]
    stack.PushFloat(val)
}

// Load int from array 
type iaload struct {NoOperandsInstruction}
func (self *iaload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    intArr := arrRef.Fields().([]int32)
    checkArrIndex(index, len(intArr))
    val := intArr[index]
    stack.PushInt(val)
}

// Load long from array 
type laload struct {NoOperandsInstruction}
func (self *laload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    longArr := arrRef.Fields().([]int64)
    checkArrIndex(index, len(longArr))
    val := longArr[index]
    stack.PushLong(val)
}

// Load short from array 
type saload struct {NoOperandsInstruction}
func (self *saload) Execute(frame *rtda.Frame) {
    stack, arrRef, index := popArrAndIndex(frame)
    shortArr := arrRef.Fields().([]int16)
    checkArrIndex(index, len(shortArr))
    val := shortArr[index]
    stack.PushInt(int32(val))
}

func popArrAndIndex(frame *rtda.Frame) (*rtda.OperandStack, *class.Obj, int) {
    stack := frame.OperandStack()
    index := int(stack.PopInt())
    arrRef := stack.PopRef()
    if arrRef == nil {
        // todo
        panic("NullPointerException")
    }
    return stack, arrRef, index
}
