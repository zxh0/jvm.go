package instructions

import "jvmgo/rtda"

// Do nothing
type nop struct {NoOperandsInstruction}
func (self *nop) Execute(thread *rtda.Thread) {
    // really do nothing
}
