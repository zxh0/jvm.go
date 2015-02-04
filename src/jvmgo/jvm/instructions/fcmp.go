package instructions

import "jvmgo/rtda"

// Compare float
type fcmpg struct {NoOperandsInstruction}
func (self *fcmpg) Execute(frame *rtda.Frame) {
    _fcmp(frame, true)
}

type fcmpl struct {NoOperandsInstruction}
func (self *fcmpl) Execute(frame *rtda.Frame) {
    _fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
    stack := frame.OperandStack()
    v2 := stack.PopFloat()
    v1 := stack.PopFloat()
    if v1 > v2 {
        stack.PushInt(1)
    } else if v1 == v2 {
        stack.PushInt(0)
    } else if v1 < v2 {
        stack.PushInt(-1)
    } else if gFlag {
        stack.PushInt(1)
    } else {
        stack.PushInt(-1)
    }
}
