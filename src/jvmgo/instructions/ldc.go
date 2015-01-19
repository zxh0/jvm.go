package instructions

import "jvmgo/rtda"

// Push item from run-time constant pool 
type ldc struct {Index8Instruction}
func (self *ldc) execute(thread *rtda.Thread) {
    // todo
    thread.CurrentFrame().Method().Class()
}

// Push item from run-time constant pool (wide index)
type ldc_w struct {Index16Instruction}
func (self *ldc_w) execute(thread *rtda.Thread) {
    // todo
}

// Push long or double from run-time constant pool (wide index) 
type ldc2_w struct {Index16Instruction}
func (self *ldc2_w) execute(thread *rtda.Thread) {
    // todo
}
