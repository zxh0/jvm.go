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
