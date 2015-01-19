package instructions

import "jvmgo/rtda"

// Throw exception or error
type athrow struct {NoOperandsInstruction}
func (self *athrow) execute(thread *rtda.Thread) {
    // todo
}
