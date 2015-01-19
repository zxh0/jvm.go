package instructions

import "jvmgo/rtda"

// Invoke a class (static) method 
type invokestatic struct {Index16Instruction}
func (self *invokestatic) execute(thread *rtda.Thread) {
    // todo
}
