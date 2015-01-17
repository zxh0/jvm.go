package instructions

import "jvmgo/rtda"

// Check whether object is of given type
type checkcast struct {
    index uint16
}

func (self *checkcast) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
}

func (self *checkcast) execute(thread *rtda.Thread) {
    // todo
}
