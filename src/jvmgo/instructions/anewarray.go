package instructions

import "jvmgo/rtda"

// Create new array of reference
type anewarray struct {
    index uint16
}

func (self *anewarray) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
}

func (self *anewarray) execute(thread *rtda.Thread) {
    // todo
}
