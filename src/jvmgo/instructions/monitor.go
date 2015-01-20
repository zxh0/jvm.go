package instructions

import "jvmgo/rtda"

// Enter monitor for object
type monitorenter struct {NoOperandsInstruction}
func (self *monitorenter) Execute(thread *rtda.Thread) {
    // todo
}

// Exit monitor for object
type monitorexit struct {NoOperandsInstruction}
func (self *monitorexit) Execute(thread *rtda.Thread) {
    // todo
}
