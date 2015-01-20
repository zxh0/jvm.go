package instructions

import "jvmgo/rtda"

type Ref *rtda.Obj

// Load reference from array
type aaload struct {NoOperandsInstruction}
func (self *aaload) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    arrRef := stack.PopRef()
    index := int(stack.PopInt())

    if arrRef == nil {
        panic("NullPointerException")
    }

    refArr := arrRef.Fields().([]Ref)   
    if index < 0 || index >= len(refArr) {
        panic("ArrayIndexOutOfBoundsException")
    }

    ref := refArr[index]
    stack.PushRef(ref)
}

// Load byte or boolean from array 
type baload struct {NoOperandsInstruction}
func (self *baload) execute(thread *rtda.Thread) {
    // todo
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
