package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Push item from run-time constant pool 
type ldc struct {Index8Instruction}
func (self *ldc) execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(self.index)
    if cInt, ok := c.(class.ConstantInt); ok {
        frame.OperandStack().PushInt(cInt.Val())
    } else if cFloat, ok := c.(class.ConstantFloat); ok {
        frame.OperandStack().PushFloat(cFloat.Val())
    }
    // todo
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
