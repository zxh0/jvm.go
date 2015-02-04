package lang

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _throwable(fillInStackTrace, "fillInStackTrace", "(I)Ljava/lang/Throwable;")
}

func _throwable(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Throwable", name, desc, method)
}

// private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(frame *rtda.Frame) {
    stack := frame.OperandStack()
    _ = stack.PopInt() // dummy
    this := stack.PopRef() // this
    stack.PushRef(this)
    // todo
}
