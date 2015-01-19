package instructions

import "jvmgo/rtda"

// Push item from run-time constant pool 
type ldc struct {Index8Instruction}
func (self *ldc) execute(thread *rtda.Thread) {
    // todo
}
