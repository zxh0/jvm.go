package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _class("desiredAssertionStatus0",   "(Ljava/lang/Class;)Z",         desiredAssertionStatus0)
    _class("getClassLoader0",           "()Ljava/lang/ClassLoader;",    getClassLoader0)
    _class("getName0",                  "()Ljava/lang/String;",         getName0)
}

func _class(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/Class", name, desc, method)
}
//java/lang/Class~getPrimitiveClass~(Ljava/lang/String;)Ljava/lang/Class;

func desiredAssertionStatus0(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // this
    stack.PushBoolean(false)
}

func getClassLoader0(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // this
    stack.PushRef(nil)
}

func getName0(frame *rtda.Frame) {
    panic("getName0")
}
