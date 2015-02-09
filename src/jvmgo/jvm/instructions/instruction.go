package instructions

import "jvmgo/jvm/rtda"

type Instruction interface {
    fetchOperands(decoder *InstructionDecoder)
    Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
    // empty
}
func (self *NoOperandsInstruction) fetchOperands(decoder *InstructionDecoder) {
    // nothing to do
}

type BranchInstruction struct {
    offset int // todo target
}
func (self *BranchInstruction) fetchOperands(decoder *InstructionDecoder) {
    self.offset = int(decoder.readInt16())
}

type Index8Instruction struct {
    index uint
}
func (self *Index8Instruction) fetchOperands(decoder *InstructionDecoder) {
    self.index = uint(decoder.readUint8())
}

type Index16Instruction struct {
    index uint
}
func (self *Index16Instruction) fetchOperands(decoder *InstructionDecoder) {
    self.index = uint(decoder.readUint16())
}
