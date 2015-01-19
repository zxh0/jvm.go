package instructions

import "jvmgo/rtda"

// Create new object
type _new struct {Index16Instruction}
func (self *_new) execute(thread *rtda.Thread) {
    // todo
}

// Create new array of reference
type anewarray struct {Index16Instruction}
func (self *anewarray) execute(thread *rtda.Thread) {
    // todo
}

// Create new array
type newarray struct {
    atype uint8
}
func (self *newarray) fetchOperands(bcr *BytecodeReader) {
    self.atype = bcr.readUint8()
}
func (self *newarray) execute(thread *rtda.Thread) {
    // todo
}
