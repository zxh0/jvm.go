package instructions

import "jvmgo/rtda"

// Fetch field from object
type getfield struct {
    index uint16
}
func (self *getfield) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
}
func (self *getfield) execute(thread *rtda.Thread) {
    //stack := thread.CurrentFrame().OperandStack()
    //ref := stack.PopRef()
    // todo
}

// Get static field from class 
type getstatic struct {
    index uint16
}
func (self *getstatic) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
}
func (self *getstatic) execute(thread *rtda.Thread) {
    // todo
}
