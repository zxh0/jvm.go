package instructions

import "jvmgo/rtda"

// Throw exception or error
type athrow struct {NoOperandsInstruction}
func (self *athrow) Execute(thread *rtda.Thread) {
    // todo
    panic("todo athrow!")
}
