package instructions

import "jvmgo/jvm/rtda"

// Do nothing
type nop struct {NoOperandsInstruction}
func (self *nop) Execute(frame *rtda.Frame) {
    // really do nothing
}
