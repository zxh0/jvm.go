package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _object(getClass, "getClass", "()Ljava/lang/Class;")
}

func _object(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Object", name, desc, method)
}

func getClass(frame *rtda.Frame) {
    stack := frame.OperandStack()
    this := stack.PopRef()
    class := this.Class().JClass()
    stack.PushRef(class)
}
