package instructions

import "jvmgo/rtda"

// Invoke a class (static) method 
type invokestatic struct {Index16Instruction}
func (self *invokestatic) execute(thread *rtda.Thread) {
    // todo
}

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations 
type invokespecial struct {Index16Instruction}
func (self *invokespecial) execute(thread *rtda.Thread) {
    // todo
}

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) execute(thread *rtda.Thread) {
    // todo
}
