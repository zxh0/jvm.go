package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Set field in object
type putfield struct {Index16Instruction}
func (self *putfield) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    ref := stack.PopRef()
    val := stack.Pop()
    if ref == nil {
        // todo NullPointerException
    }
    
    cp := frame.Method().Class().ConstantPool()
    cFieldRef := cp.GetConstant(self.index).(class.ConstantFieldref)
    field := cFieldRef.Field()

    field.PutValue(ref, val)
}

// Set static field in class
type putstatic struct {Index16Instruction}
func (self *putstatic) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    val := stack.Pop()
    
    cp := frame.Method().Class().ConstantPool()
    cFieldRef := cp.GetConstant(self.index).(class.ConstantFieldref)
    field := cFieldRef.Field()

    field.PutStaticValue(val)
}
