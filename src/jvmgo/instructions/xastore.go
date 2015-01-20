package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
)

// Store into reference array 
type aastore struct {NoOperandsInstruction}
func (self *aastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into byte or boolean array 
type bastore struct {NoOperandsInstruction}
func (self *bastore) execute(thread *rtda.Thread) {
    arrRef, index, val := popOperands(thread)
    byteArr := arrRef.Fields().([]int8)
    index = checkIndex(index, len(byteArr))
    byteArr[index] = val.(int8)
}

// Store into char array 
type castore struct {NoOperandsInstruction}
func (self *castore) execute(thread *rtda.Thread) {
    // todo
}

// Store into double array 
type dastore struct {NoOperandsInstruction}
func (self *dastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into float array 
type fastore struct {NoOperandsInstruction}
func (self *fastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into int array 
type iastore struct {NoOperandsInstruction}
func (self *iastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into long array 
type lastore struct {NoOperandsInstruction}
func (self *lastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into short array 
type sastore struct {NoOperandsInstruction}
func (self *sastore) execute(thread *rtda.Thread) {
    // todo
}

func popOperands(thread *rtda.Thread) (*rtda.Obj, int, Any) {
    stack := thread.CurrentFrame().OperandStack()
    arrRef := stack.PopRef()
    index := int(stack.PopInt())
    val := stack.Pop()
    if arrRef == nil {
        // todo
        panic("NullPointerException")
    }
    return arrRef, index, val
}
