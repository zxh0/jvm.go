package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
)

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
    execute(thread *rtda.Thread)
}

type NoOperandsInstruction struct {
    // empty
}
func (self *NoOperandsInstruction) fetchOperands(bcr *BytecodeReader) {
    // nothing to do
}

type BranchInstruction struct {
    branch int16
}
func (self *BranchInstruction) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}

type Index8Instruction struct {
    index uint
}
func (self *Index8Instruction) fetchOperands(bcr *BytecodeReader) {
    self.index = uint(bcr.readUint8())
}

type Index16Instruction struct {
    index uint16
}
func (self *Index16Instruction) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
}

// todo: move to any.go?
func isLongOrDouble(x Any) (bool) {
    if _, ok := x.(int64); ok {
        return true
    }
    if _, ok := x.(float64); ok {
        return true
    }
    return false
}
