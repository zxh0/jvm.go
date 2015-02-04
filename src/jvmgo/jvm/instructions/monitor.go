package instructions

import "jvmgo/jvm/rtda"

// Enter monitor for object
type monitorenter struct {NoOperandsInstruction}
func (self *monitorenter) Execute(frame *rtda.Frame) {
    ref := frame.OperandStack().PopRef()
    if ref != nil {
        
    }
    // todo
    //panic("monitorenter")
}

// Exit monitor for object
type monitorexit struct {NoOperandsInstruction}
func (self *monitorexit) Execute(frame *rtda.Frame) {
    ref := frame.OperandStack().PopRef()
    if ref != nil {
        
    }
    // todo
    //panic("monitorexit")
}
