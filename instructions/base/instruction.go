package base

import (
	"github.com/zxh0/jvm.go/rtda"
)

type Instruction interface {
	FetchOperands(reader *CodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
	// empty
}

func (instr *NoOperandsInstruction) FetchOperands(reader *CodeReader) {
	// nothing to do
}

type BranchInstruction struct {
	Offset int // todo target
}

func (instr *BranchInstruction) FetchOperands(reader *CodeReader) {
	instr.Offset = int(reader.ReadInt16())
}

// only used by instruction decoder
func (instr *BranchInstruction) GetOffset() int       { return instr.Offset }
func (instr *BranchInstruction) SetOffset(offset int) { instr.Offset = offset }

type Index8Instruction struct {
	Index uint
}

func (instr *Index8Instruction) FetchOperands(reader *CodeReader) {
	instr.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (instr *Index16Instruction) FetchOperands(reader *CodeReader) {
	instr.Index = uint(reader.ReadUint16())
}
