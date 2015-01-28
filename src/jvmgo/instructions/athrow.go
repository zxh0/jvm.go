package instructions

import "jvmgo/rtda"

// Throw exception or error
type athrow struct {NoOperandsInstruction}
func (self *athrow) Execute(thread *rtda.Thread) {
    currentFrame := thread.CurrentFrame()
    e := currentFrame.OperandStack().PopRef()
    if e == nil {
        // todo NPE
        panic("athrow NPE")
    }

    currentFrame.Method()

    // todo
    panic("todo athrow!")
}
