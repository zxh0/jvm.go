package instructions

import "jvmgo/rtda"

// Load reference from array
type aaload struct {}
func (self *aaload) fetchOperands(bcr *BytecodeReader) {}
func (self *aaload) execute(thread *rtda.Thread) {
    // todo
}

// Load byte or boolean from array 
type baload struct {}
func (self *baload) fetchOperands(bcr *BytecodeReader) {}
func (self *baload) execute(thread *rtda.Thread) {
    // todo
}

// Load char from array 
type caload struct {}
func (self *caload) fetchOperands(bcr *BytecodeReader) {}
func (self *caload) execute(thread *rtda.Thread) {
    // todo
}

// Load double from array 
type daload struct {}
func (self *daload) fetchOperands(bcr *BytecodeReader) {}
func (self *daload) execute(thread *rtda.Thread) {
    // todo
}

// Load float from array 
type faload struct {}
func (self *faload) fetchOperands(bcr *BytecodeReader) {}
func (self *faload) execute(thread *rtda.Thread) {
    // todo
}
