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
