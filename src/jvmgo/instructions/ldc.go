package instructions

import "jvmgo/rtda"

// Push item from run-time constant pool 
type ldc struct {Index8Instruction}
func (self *ldc) execute(thread *rtda.Thread) {
    // todo
}

// Push item from run-time constant pool (wide index)
type ldc_w struct {Index16Instruction}
func (self *ldc_w) execute(thread *rtda.Thread) {
    // todo
}
