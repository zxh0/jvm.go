package instructions

import "jvmgo/rtda"

// Create new object
type _new struct {Index16Instruction}
func (self *_new) execute(thread *rtda.Thread) {
    // todo
}

// Create new array of reference
type anewarray struct {Index16Instruction}
func (self *anewarray) execute(thread *rtda.Thread) {
    // todo
}
