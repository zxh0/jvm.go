package control

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

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
type LookupSwitch struct { // TODO
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (instr *LookupSwitch) FetchOperands(reader *base.CodeReader) {
	reader.SkipPadding()
	instr.defaultOffset = reader.ReadInt32()
	instr.npairs = reader.ReadInt32()
	instr.matchOffsets = reader.ReadInt32s(instr.npairs * 2)
}

func (instr *LookupSwitch) Execute(frame *rtda.Frame) {
	key := frame.PopInt()
	for i := int32(0); i < instr.npairs*2; i += 2 {
		if instr.matchOffsets[i] == key {
			offset := instr.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(instr.defaultOffset))
}
