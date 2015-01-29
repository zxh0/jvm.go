package instructions

import (
    //"fmt"
    "jvmgo/rtda"
)

// Throw exception or error
type athrow struct {NoOperandsInstruction}
func (self *athrow) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()

    ex := frame.OperandStack().PopRef()
    if ex == nil {
        // todo NPE
        panic("athrow NPE")
    }

    handler := frame.Method().FindExceptionHandler(ex.Class(), thread.PC())
    if handler != nil {
        stack.PushRef(ex)
        frame.SetNextPC(handler.HandlerPc())
        return
    }

    // todo
    panic("todo athrow!")
}
