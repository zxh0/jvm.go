package instructions

import (
    //"fmt"
    "jvmgo/rtda"
)

// Access jump table by index and jump
type tableswitch struct {
    defaultOffset   int32
    low             int32
    high            int32
    jumpOffsets     []int32
}

func (self *tableswitch) fetchOperands(bcr *BytecodeReader) {
    // skip padding
    for bcr.pc % 4 != 0 {
        bcr.readUint8()
    }
    self.defaultOffset = bcr.readInt32()
    self.low = bcr.readInt32()
    self.high = bcr.readInt32()
    jumpOffsetsCount := self.high - self.low + 1
    self.jumpOffsets = make([]int32, jumpOffsetsCount)
    for i := range self.jumpOffsets {
        self.jumpOffsets[i] = bcr.readInt32()
    }
}

func (self *tableswitch) Execute(thread *rtda.Thread) {
    index := thread.CurrentFrame().OperandStack().PopInt()

    var offset int32
    if index >= self.low && index <= self.high {
        offset = self.jumpOffsets[index - self.low]
    } else {
        offset = self.defaultOffset
    }

    branch(thread, int(offset))
}
