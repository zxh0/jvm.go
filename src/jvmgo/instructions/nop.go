package instructions

import "jvmgo/rtda"

// Do nothing
type nop struct {NoOperandsInstruction}
func (self *nop) execute(thread *rtda.Thread) {
    // really do nothing
}
