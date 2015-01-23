package instructions

import (
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Check whether object is of given type
type checkcast struct {Index16Instruction}
func (self *checkcast) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    ref := stack.PopRef()
    stack.PushRef(ref)

    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(rtc.ConstantClass)
    class := cClass.Class()
    if class.NotInitialized() {
        // todo init class
    }

    // todo
    if !_instanceof(ref, class) {
        // todo ClassCastException
    }
}
