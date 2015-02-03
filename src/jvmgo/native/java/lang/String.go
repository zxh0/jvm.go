package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _string(intern, "intern", "()Ljava/lang/String;")
}

func _string(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/String", name, desc, method)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
    //panic("intern!!")
    stack := frame.OperandStack()
    this := stack.PopRef()
    stack.PushRef(this)
    // todo
}
