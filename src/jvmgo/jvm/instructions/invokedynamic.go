package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/jvm/rtda"
    //"jvmgo/jvm/rtda/class"
)

// Invoke dynamic method
type invokedynamic struct {
    index uint16
    // 0
    // 0
}
func (self *invokedynamic) fetchOperands(decoder *InstructionDecoder) {
    self.index = decoder.readUint16()
    decoder.readUint8() // must be 0
    decoder.readUint8() // must be 0
}
func (self *invokedynamic) Execute(frame *rtda.Frame) {
    // todo
    panic("todo invokedynamic")
}
