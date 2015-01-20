package instructions

import "jvmgo/rtda"

//type Ref *rtda.Obj

// Load reference from array
type aaload struct {NoOperandsInstruction}
func (self *aaload) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    arrRef := popAndCheckArrRef(stack)
    index := int(stack.PopInt())

    refArr := arrRef.Fields().([]*rtda.Obj)
    ref := refArr[checkIndex(index, len(refArr))]
    stack.PushRef(ref)
}

// Load byte or boolean from array 
type baload struct {NoOperandsInstruction}
func (self *baload) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    arrRef := popAndCheckArrRef(stack)
    index := int(stack.PopInt())

    byteArr := arrRef.Fields().([]int8)
    val := byteArr[checkIndex(index, len(byteArr))]
    stack.PushInt(int32(val))
}

// Load char from array 
type caload struct {NoOperandsInstruction}
func (self *caload) execute(thread *rtda.Thread) {
    // todo
}

// Load double from array 
type daload struct {NoOperandsInstruction}
func (self *daload) execute(thread *rtda.Thread) {
    // todo
}

// Load float from array 
type faload struct {NoOperandsInstruction}
func (self *faload) execute(thread *rtda.Thread) {
    // todo
}

// Load int from array 
type iaload struct {NoOperandsInstruction}
func (self *iaload) execute(thread *rtda.Thread) {
    // todo
}

// Load long from array 
type laload struct {NoOperandsInstruction}
func (self *laload) execute(thread *rtda.Thread) {
    // todo
}

// Load short from array 
type saload struct {NoOperandsInstruction}
func (self *saload) execute(thread *rtda.Thread) {
    // todo
}

func popAndCheckArrRef(stack *rtda.OperandStack) (*rtda.Obj) {
    arrRef := stack.PopRef()
    if arrRef == nil {
        // todo
        panic("NullPointerException")
    }
    return arrRef
}

func checkIndex(index, len int) (int) {
    if index < 0 || index >= len {
        panic("ArrayIndexOutOfBoundsException")
    }
    return index
}
