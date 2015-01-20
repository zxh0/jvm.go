package instructions

import "jvmgo/rtda"

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
    Execute(thread *rtda.Thread)
}

type NoOperandsInstruction struct {
    // empty
}
func (self *NoOperandsInstruction) fetchOperands(bcr *BytecodeReader) {
    // nothing to do
}

type BranchInstruction struct {
    branch int
}
func (self *BranchInstruction) fetchOperands(bcr *BytecodeReader) {
    self.branch = int(bcr.readInt16())
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
