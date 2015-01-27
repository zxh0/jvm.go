package instructions

import (
    //"fmt"
    "jvmgo/rtda"
)

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
type tableswitch struct {
    defaultOffset   int32
    low             int32
    high            int32
    jumpOffsets     []int32
}

func (self *tableswitch) fetchOperands(bcr *BytecodeReader) {
    for bcr.pc % 4 != 0 {
        // skip padding
        bcr.readUint8()
    }
    self.defaultOffset = bcr.readInt32()
    self.low = bcr.readInt32()
    self.high = bcr.readInt32()
    jumpOffsetsCount := self.high - self.low + 1
    self.jumpOffsets = bcr.readInt32s(jumpOffsetsCount)
}

func (self *tableswitch) Execute(thread *rtda.Thread) {
    index := thread.CurrentFrame().OperandStack().PopInt()

    var offset int
    if index >= self.low && index <= self.high {
        offset = int(self.jumpOffsets[index - self.low])
    } else {
        offset = int(self.defaultOffset)
    }

    branch(thread, offset)
}
