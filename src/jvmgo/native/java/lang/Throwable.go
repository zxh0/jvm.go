package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _throwable("fillInStackTrace",  "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func _throwable(name, desc string, method Any) {
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
