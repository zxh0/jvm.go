package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _object("getClass", "()Ljava/lang/Class;",  getClass)
}

func _object(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/Object", name, desc, method)
}

func getClass(stack *rtda.OperandStack) {
    this := stack.PopRef()
    class := this.Class().JClass()
    stack.PushRef(class)
}
