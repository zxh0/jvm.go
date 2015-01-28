package instructions

import "jvmgo/rtda"

// Throw exception or error
type athrow struct {NoOperandsInstruction}
func (self *athrow) Execute(thread *rtda.Thread) {
    currentFrame := thread.CurrentFrame()
    ex := currentFrame.OperandStack().PopRef()
    if ex == nil {
        // todo NPE
        panic("athrow NPE")
    }

    handler := currentFrame.Method().FindExceptionHandler(ex.Class())
    if handler != nil {
        panic("here!!!")
    }

    // todo
    panic("todo athrow!")
}
