package instructions

import (
    //"fmt"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Determine if object is of given type
type instanceof struct {Index16Instruction}
func (self *instanceof) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    ref := stack.PopRef()

    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
    class := cClass.Class()

    if class.InitializationNotStarted() {
        // todo init class
        panic("class not initialized!" + class.Name())
    }

    if ref == nil {
        stack.PushInt(0)
    } else if ref.IsInstanceOf(class) {
        stack.PushInt(1)
    } else {
        stack.PushInt(0)
    }
}
