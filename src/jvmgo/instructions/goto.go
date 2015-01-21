package instructions

import "jvmgo/rtda"

// Branch always
type goto_ struct {BranchInstruction}
func (self *goto_) Execute(thread *rtda.Thread) {
    branch(thread, self.offset)
}

// Branch always (wide index) 
type goto_w struct {
    offset int
}
func (self *goto_w) fetchOperands(bcr *BytecodeReader) {
    self.offset = int(bcr.readInt32())
}
func (self *goto_w) Execute(thread *rtda.Thread) {
    branch(thread, self.offset)
}
