package instructions

import "jvmgo/rtda"

// Create new object
type _new struct {Index16Instruction}
func (self *_new) Execute(thread *rtda.Thread) {
    // todo
}

// Create new array of reference
type anewarray struct {Index16Instruction}
func (self *anewarray) Execute(thread *rtda.Thread) {
    // todo
}

// Create new array
type newarray struct {
    atype uint8
}
func (self *newarray) fetchOperands(bcr *BytecodeReader) {
    self.atype = bcr.readUint8()
}
func (self *newarray) Execute(thread *rtda.Thread) {
    // todo
}

// Create new multidimensional array
type multianewarray struct {
    index       uint16
    dimensions  uint8
}
func (self *multianewarray) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    self.dimensions = bcr.readUint8()
}
func (self *multianewarray) Execute(thread *rtda.Thread) {
    // todo
}
