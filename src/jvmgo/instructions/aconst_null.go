package instructions

import "jvmgo/rtda"

// Push null
type aconst_null struct {NoOperandsInstruction}
func (self *aconst_null) execute(thread *rtda.Thread) {
    thread.CurrentFrame().OperandStack().PushNull()
}
