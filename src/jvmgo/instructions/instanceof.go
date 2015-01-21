package instructions

import (
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Determine if object is of given type
type instanceof struct {Index16Instruction}
func (self *instanceof) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    ref := stack.PopRef()

    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(rtc.ConstantClass)
    class := cClass.Class()

    if !class.IsInitialized() {
        // todo init class
    }

    // todo
    if _instanceof(ref, class) {
        stack.PushInt(1)
    } else {
        stack.PushInt(0)
    }
}

func _instanceof(ref *rtc.Obj, class *rtc.Class) (bool) {
    // todo
    return false
}
