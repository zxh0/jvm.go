package instructions

import "jvmgo/rtda"

// Determine if object is of given type
type instanceof struct {
    index uint16
}

func (self *instanceof) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
}

func (self *instanceof) execute(thread *rtda.Thread) {
    // todo
}
