package instructions

import "jvmgo/rtda"

// Fetch field from object
type getfield struct {Index16Instruction}
func (self *getfield) execute(thread *rtda.Thread) {
    //stack := thread.CurrentFrame().OperandStack()
    //ref := stack.PopRef()
    // todo
}

// Get static field from class 
type getstatic struct {Index16Instruction}
func (self *getstatic) execute(thread *rtda.Thread) {
    // todo
}
