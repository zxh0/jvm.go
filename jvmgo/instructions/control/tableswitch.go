package instructions

import (
	//"fmt"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
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
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *tableswitch) fetchOperands(decoder *InstructionDecoder) {
	for decoder.pc%4 != 0 {
		// skip padding
		decoder.readUint8()
	}
	self.defaultOffset = decoder.readInt32()
	self.low = decoder.readInt32()
	self.high = decoder.readInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = decoder.readInt32s(jumpOffsetsCount)
}

func (self *tableswitch) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	branch(frame, offset)
}
