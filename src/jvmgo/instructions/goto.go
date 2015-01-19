package instructions

import "jvmgo/rtda"

// Branch always
type _goto struct {
    branch int16
}
func (self *_goto) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *_goto) execute(thread *rtda.Thread) {
    // todo
}

// Branch always (wide index) 
type goto_w struct {
    branch int32
}
func (self *goto_w) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt32()
}
func (self *goto_w) execute(thread *rtda.Thread) {
    // todo
}
