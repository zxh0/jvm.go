package instructions

import "jvmgo/rtda"

// Enter monitor for object
type monitorenter struct {NoOperandsInstruction}
func (self *monitorenter) execute(thread *rtda.Thread) {
    // todo
}

// Exit monitor for object
type monitorexit struct {NoOperandsInstruction}
func (self *monitorexit) execute(thread *rtda.Thread) {
    // todo
}
