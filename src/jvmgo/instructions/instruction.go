package instructions

import "jvmgo/rtda"

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
    // todo Execute(frame *rtda.Frame)
    Execute(thread *rtda.Thread)
}

type NoOperandsInstruction struct {
    // empty
}
func (self *NoOperandsInstruction) fetchOperands(bcr *BytecodeReader) {
    // nothing to do
}

type BranchInstruction struct {
    offset int
}
func (self *BranchInstruction) fetchOperands(bcr *BytecodeReader) {
    self.offset = int(bcr.readInt16())
}

type Index8Instruction struct {
    index uint
}
func (self *Index8Instruction) fetchOperands(bcr *BytecodeReader) {
    self.index = uint(bcr.readUint8())
}

type Index16Instruction struct {
    index uint
}
func (self *Index16Instruction) fetchOperands(bcr *BytecodeReader) {
    self.index = uint(bcr.readUint16())
}
