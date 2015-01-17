package instructions

import "jvmgo/rtda"

// Create new array of reference
type anewarray struct {
    index uint16
}

func (self *anewarray) fetchOperands(bcr *BytecodeReader) {
    indexbyte1 := uint16(bcr.readUint8())
    indexbyte2 := uint16(bcr.readUint8())
    self.index = (indexbyte1 << 8) | indexbyte2
}

func (self *anewarray) execute(thread *rtda.Thread) {
    // todo
}
