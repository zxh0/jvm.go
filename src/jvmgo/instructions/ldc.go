package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Push item from run-time constant pool 
type ldc struct {Index8Instruction}
func (self *ldc) execute(thread *rtda.Thread) {
    _ldc(thread, self.index)
}

// Push item from run-time constant pool (wide index)
type ldc_w struct {Index16Instruction}
func (self *ldc_w) execute(thread *rtda.Thread) {
    _ldc(thread, self.index)
}

func _ldc(thread *rtda.Thread, index uint) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(index)

    if cInt, ok := c.(class.ConstantInt); ok {
        stack.PushInt(cInt.Val())
    } else if cFloat, ok := c.(class.ConstantFloat); ok {
        stack.PushFloat(cFloat.Val())
    }
    // todo
    // ref to String
    // ref to Class
    // ref to MethodType or MethodHandle
}

// Push long or double from run-time constant pool (wide index) 
type ldc2_w struct {Index16Instruction}
func (self *ldc2_w) execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(self.index)

    if cLong, ok := c.(class.ConstantLong); ok {
        stack.PushLong(cLong.Val())
    } else if cDouble, ok := c.(class.ConstantDouble); ok {
        stack.PushDouble(cDouble.Val())
    } else {
        // todo
    }
}
