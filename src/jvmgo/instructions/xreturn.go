package instructions

import "jvmgo/rtda"

// Return reference from method 
type areturn struct {}
func (self *areturn) fetchOperands(bcr *BytecodeReader) {}
func (self *areturn) execute(thread *rtda.Thread) {
    // todo
}

// Return double from method 
type dreturn struct {}
func (self *dreturn) fetchOperands(bcr *BytecodeReader) {}
func (self *dreturn) execute(thread *rtda.Thread) {
    // todo
}

// Return float from method 
type freturn struct {}
func (self *freturn) fetchOperands(bcr *BytecodeReader) {}
func (self *freturn) execute(thread *rtda.Thread) {
    // todo
}

// Return int from method 
type ireturn struct {}
func (self *ireturn) fetchOperands(bcr *BytecodeReader) {}
func (self *ireturn) execute(thread *rtda.Thread) {
    // todo
}

// Return double from method 
type lreturn struct {}
func (self *lreturn) fetchOperands(bcr *BytecodeReader) {}
func (self *lreturn) execute(thread *rtda.Thread) {
    // todo
}
