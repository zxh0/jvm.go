package instructions

import "jvmgo/rtda"

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type lookupswitch struct {
    defaultOffset   int32
    npairs          int32
    matchOffsets    []int32
}

func (self *lookupswitch) fetchOperands(bcr *BytecodeReader) {
    for bcr.pc % 4 != 0 {
        // skip padding
        bcr.readUint8()
    }
    self.defaultOffset = bcr.readInt32()
    self.npairs = bcr.readInt32()
    self.matchOffsets = bcr.readInt32s(self.npairs * 2)
}

func (self *lookupswitch) Execute(frame *rtda.Frame) {
    key := frame.OperandStack().PopInt()
    for i := int32(0); i < self.npairs * 2; i+=2 {
        if self.matchOffsets[i] == key {
            offset := self.matchOffsets[i + 1]
            branch(frame, int(offset))
            return
        }
    }
    branch(frame, int(self.defaultOffset))
}
