package instructions

import "jvmgo/rtda"

// Branch always
type _goto struct {BranchInstruction}
func (self *_goto) Execute(thread *rtda.Thread) {
    // todo
}

// Branch always (wide index) 
type goto_w struct {
    branch int32
}
func (self *goto_w) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt32()
}
func (self *goto_w) Execute(thread *rtda.Thread) {
    // todo
}
