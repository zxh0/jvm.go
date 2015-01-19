package instructions

import "jvmgo/rtda"

// Branch always
type _goto struct {
    index int16
}
func (self *_goto) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readInt16()
}
func (self *_goto) execute(thread *rtda.Thread) {
    // todo
}

// Branch always (wide index) 
type goto_w struct {
    index int16
}
func (self *goto_w) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readInt16()
}
func (self *goto_w) execute(thread *rtda.Thread) {
    // todo
}
