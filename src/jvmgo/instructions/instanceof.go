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

    if class.NotInitialized() {
        // todo init class
        panic("class not initialized!" + class.Name())
    }

    // todo
    if _instanceof(ref, class) {
        stack.PushInt(1)
    } else {
        stack.PushInt(0)
    }
}

func _instanceof(ref *rtc.Obj, class *rtc.Class) (bool) {
    if ref == nil {
        return false
    } else {
        return ref.Class() == class
    }
}
