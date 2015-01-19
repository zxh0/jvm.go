package instructions

import "jvmgo/rtda"

// Return void from method 
type _return struct {NoOperandsInstruction}
func (self *_return) execute(thread *rtda.Thread) {
    // todo
}

// Return reference from method 
type areturn struct {NoOperandsInstruction}
func (self *areturn) execute(thread *rtda.Thread) {
    // todo
}

// Return double from method 
type dreturn struct {NoOperandsInstruction}
func (self *dreturn) execute(thread *rtda.Thread) {
    // todo
}

// Return float from method 
type freturn struct {NoOperandsInstruction}
func (self *freturn) execute(thread *rtda.Thread) {
    // todo
}

// Return int from method 
type ireturn struct {NoOperandsInstruction}
func (self *ireturn) execute(thread *rtda.Thread) {
    // todo
}

// Return double from method 
type lreturn struct {NoOperandsInstruction}
func (self *lreturn) execute(thread *rtda.Thread) {
    // todo
}
