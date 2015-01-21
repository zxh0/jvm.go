package instructions

import "jvmgo/rtda"

// Return void from method 
type return_ struct {NoOperandsInstruction}
func (self *return_) Execute(thread *rtda.Thread) {
    thread.PopFrame()
}

// Return reference from method 
type areturn struct {NoOperandsInstruction}
func (self *areturn) Execute(thread *rtda.Thread) {
    currentFrame := thread.PopFrame()
    invokerFrame := thread.TopFrame()
    ref := currentFrame.OperandStack().PopRef()
    invokerFrame.OperandStack().PushRef(ref)
}

// Return double from method 
type dreturn struct {NoOperandsInstruction}
func (self *dreturn) Execute(thread *rtda.Thread) {
    // todo
}

// Return float from method 
type freturn struct {NoOperandsInstruction}
func (self *freturn) Execute(thread *rtda.Thread) {
    // todo
}

// Return int from method 
type ireturn struct {NoOperandsInstruction}
func (self *ireturn) Execute(thread *rtda.Thread) {
    // todo
}

// Return double from method 
type lreturn struct {NoOperandsInstruction}
func (self *lreturn) Execute(thread *rtda.Thread) {
    // todo
}
