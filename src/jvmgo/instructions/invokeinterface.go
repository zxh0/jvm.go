package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/rtda"
    //"jvmgo/rtda/class"
)

// Invoke interface method
type invokeinterface struct {
    index   uint16
    count   uint8
    // 0
}
func (self *invokeinterface) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    self.count = bcr.readUint8()
    bcr.readUint8() // must be 0
}
func (self *invokeinterface) Execute(frame *rtda.Frame) {
    // todo
    panic("todo invokeinterface")
}
