package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/rtda"
    //"jvmgo/rtda/class"
)

// Invoke dynamic method
type invokedynamic struct {
    index uint16
    // 0
    // 0
}
func (self *invokedynamic) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    bcr.readUint8() // must be 0
    bcr.readUint8() // must be 0
}
func (self *invokedynamic) Execute(thread *rtda.Thread) {
    // todo
    panic("todo invokedynamic")
}
