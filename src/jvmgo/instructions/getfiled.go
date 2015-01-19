package instructions

import "jvmgo/rtda"

// Store float into local variable 
type getfield struct {
    index uint16
}
func (self *getfield) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
}
func (self *getfield) execute(thread *rtda.Thread) {
    //stack := thread.CurrentFrame().OperandStack()
    //ref := stack.PopRef()
}
