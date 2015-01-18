package instructions

import "jvmgo/rtda"

// Store into reference array 
type aastore struct {}
func (self *aastore) fetchOperands(bcr *BytecodeReader) {}
func (self *aastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into byte or boolean array 
type bastore struct {}
func (self *bastore) fetchOperands(bcr *BytecodeReader) {}
func (self *bastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into char array 
type castore struct {}
func (self *castore) fetchOperands(bcr *BytecodeReader) {}
func (self *castore) execute(thread *rtda.Thread) {
    // todo
}

// Store into double array 
type dastore struct {}
func (self *dastore) fetchOperands(bcr *BytecodeReader) {}
func (self *dastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into float array 
type fastore struct {}
func (self *fastore) fetchOperands(bcr *BytecodeReader) {}
func (self *fastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into int array 
type iastore struct {}
func (self *iastore) fetchOperands(bcr *BytecodeReader) {}
func (self *iastore) execute(thread *rtda.Thread) {
    // todo
}

// Store into long array 
type lastore struct {}
func (self *lastore) fetchOperands(bcr *BytecodeReader) {}
func (self *lastore) execute(thread *rtda.Thread) {
    // todo
}
