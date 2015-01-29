package instructions

import (
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Check whether object is of given type
type checkcast struct {Index16Instruction}
func (self *checkcast) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    ref := stack.PopRef()
    stack.PushRef(ref)

    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
    class := cClass.Class()
    if class.InitializationNotStarted() {
        // todo init class
        panic("class not initialized!")
    }

    // todo
    if !_instanceof(ref, class) {
        // todo ClassCastException
        panic("ClassCastException")
    }
}
