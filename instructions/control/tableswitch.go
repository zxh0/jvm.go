package control

import (
	//"fmt"
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
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
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (instr *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	instr.defaultOffset = reader.ReadInt32()
	instr.low = reader.ReadInt32()
	instr.high = reader.ReadInt32()
	jumpOffsetsCount := instr.high - instr.low + 1
	instr.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (instr *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= instr.low && index <= instr.high {
		offset = int(instr.jumpOffsets[index-instr.low])
	} else {
		offset = int(instr.defaultOffset)
	}

	base.Branch(frame, offset)
}
