package instructions

import "jvmgo/rtda"

// Load reference from array
type aaload struct {}
func (self *aaload) fetchOperands(bcr *BytecodeReader) {}

func (self *aaload) execute(thread *rtda.Thread) {
    // todo
}
