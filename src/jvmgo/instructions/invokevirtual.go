package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/rtda"
    //"jvmgo/rtda/class"
)

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) Execute(thread *rtda.Thread) {
    // todo
    panic("todo invokevirtual")
}
