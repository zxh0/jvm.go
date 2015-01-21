package instructions

import "jvmgo/rtda"

// Branch always
type _goto struct {BranchInstruction}
func (self *_goto) Execute(thread *rtda.Thread) {
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
