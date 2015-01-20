package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Fetch field from object
type getfield struct {Index16Instruction}
func (self *getfield) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    ref := stack.PopRef()
    if ref == nil {
        // todo NullPointerException
    }

    cp := frame.Method().Class().ConstantPool()
    cFieldRef := cp.GetConstant(self.index).(class.ConstantFieldref)
    field := cFieldRef.Field()

    val := field.GetValue(ref)
    stack.Push(val)
}

// Get static field from class 
type getstatic struct {Index16Instruction}
func (self *getstatic) Execute(thread *rtda.Thread) {
    // todo
}
