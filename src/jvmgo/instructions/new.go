package instructions

import (
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Create new object
type new_ struct {Index16Instruction}
func (self *new_) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
    class := cClass.Class()

    if class.InitializationNotStarted() {
        thread := frame.Thread()
        frame.SetNextPC(thread.PC()) // undo new
        initClass(class, thread)
    } else {
        ref := class.NewObj()
        frame.OperandStack().PushRef(ref)
    }
}
