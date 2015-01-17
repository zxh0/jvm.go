package instructions

import "jvmgo/rtda"

// Check whether object is of given type
type checkcast struct {
    index uint16
}

func (self *checkcast) fetchOperands(bcr *BytecodeReader) {
    indexbyte1 := uint16(bcr.readUint8())
    indexbyte2 := uint16(bcr.readUint8())
    self.index = (indexbyte1 << 8) | indexbyte2
}

func (self *checkcast) execute(thread *rtda.Thread) {
    // todo
}
