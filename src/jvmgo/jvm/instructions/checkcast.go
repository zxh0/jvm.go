package instructions

import (
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

// Check whether object is of given type
type checkcast struct {Index16Instruction}
func (self *checkcast) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    ref := stack.PopRef()
    stack.PushRef(ref)

    if ref == nil {
        return
    }

    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
    class := cClass.Class()
    if class.InitializationNotStarted() {
        // todo init class
        panic("class not initialized!")
    }

    // todo
    if !ref.IsInstanceOf(class) {
        // todo ClassCastException
        panic("ClassCastException")
    }
}
