package lang

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _string(intern, "intern", "()Ljava/lang/String;")
}

func _string(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/String", name, desc, method)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame, x int) {
    vars := frame.LocalVars()
    str := vars.GetRef(0) // this

    chars := rtda.JStringChars(str)
    internedStr := rtda.InternString(chars, str)
    stack := frame.OperandStack()
    stack.PushRef(internedStr)
}
