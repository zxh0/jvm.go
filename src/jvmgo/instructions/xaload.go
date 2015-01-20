package instructions

import "jvmgo/rtda"

//type Ref *rtda.Obj

// Load reference from array
type aaload struct {NoOperandsInstruction}
func (self *aaload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    refArr := arrRef.Fields().([]*rtda.Obj)
    checkArrIndex(index, len(refArr))
    ref := refArr[index]
    stack.PushRef(ref)
}

// Load byte or boolean from array 
type baload struct {NoOperandsInstruction}
func (self *baload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    byteArr := arrRef.Fields().([]int8)
    checkArrIndex(index, len(byteArr))
    val := byteArr[index]
    stack.PushInt(int32(val))
}

// Load char from array 
type caload struct {NoOperandsInstruction}
func (self *caload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    charArr := arrRef.Fields().([]uint16)
    checkArrIndex(index, len(charArr))
    val := charArr[index]
    stack.PushInt(int32(val))
}

// Load double from array 
type daload struct {NoOperandsInstruction}
func (self *daload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    doubleArr := arrRef.Fields().([]float64)
    checkArrIndex(index, len(doubleArr))
    val := doubleArr[index]
    stack.PushDouble(val)
}

// Load float from array 
type faload struct {NoOperandsInstruction}
func (self *faload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    floatArr := arrRef.Fields().([]float32)
    checkArrIndex(index, len(floatArr))
    val := floatArr[index]
    stack.PushFloat(val)
}

// Load int from array 
type iaload struct {NoOperandsInstruction}
func (self *iaload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    intArr := arrRef.Fields().([]int32)
    checkArrIndex(index, len(intArr))
    val := intArr[index]
    stack.PushInt(val)
}

// Load long from array 
type laload struct {NoOperandsInstruction}
func (self *laload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    longArr := arrRef.Fields().([]int64)
    checkArrIndex(index, len(longArr))
    val := longArr[index]
    stack.PushLong(val)
}

// Load short from array 
type saload struct {NoOperandsInstruction}
func (self *saload) execute(thread *rtda.Thread) {
    stack, arrRef, index := popArrAndIndex(thread)
    shortArr := arrRef.Fields().([]int16)
    checkArrIndex(index, len(shortArr))
    val := shortArr[index]
    stack.PushInt(int32(val))
}

func popArrAndIndex(thread *rtda.Thread) (*rtda.OperandStack, *rtda.Obj, int) {
    stack := thread.CurrentFrame().OperandStack()
    arrRef := stack.PopRef()
    index := int(stack.PopInt())
    if arrRef == nil {
        // todo
        panic("NullPointerException")
    }
    return stack, arrRef, index
}
