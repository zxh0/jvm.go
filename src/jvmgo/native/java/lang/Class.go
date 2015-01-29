package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    jlClass("getName0",                 "()Ljava/lang/String;",         getName0)
    jlClass("getClassLoader0",          "()Ljava/lang/ClassLoader;",    getClassLoader0)
    jlClass("desiredAssertionStatus0",  "(Ljava/lang/Class;)Z",         desiredAssertionStatus0)
}

func jlClass(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/Class", name, desc, method)
}

func getName0(stack *rtda.OperandStack) {
    panic("getName0")
}

func getClassLoader0(stack *rtda.OperandStack) {
    // todo
    _ = stack.PopRef() // this
    stack.PushRef(nil)
}

func desiredAssertionStatus0(stack *rtda.OperandStack) {
    // todo
    _ = stack.PopRef() // this
    stack.PushBoolean(false)
}
