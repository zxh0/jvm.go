package native

import (
    "fmt"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
    _ "jvmgo/native/java/lang"
)

// register native methods
func init() {
    rtc.SetRegisterNatives(registerNatives)
    jlThrowable ("fillInStackTrace",        "(I)Ljava/lang/Throwable;",     fillInStackTrace)
    // hack!
    rtc.RegisterNativeMethod("jvmgo/SystemOut", "println", "(Ljava/lang/String;)V", jvmgo_SystemOut_println)
}

func registerNatives(operandStack *rtda.OperandStack) {
    // todo
}

func jlThrowable(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/Throwable", name, desc, method)
}

// java.lang.Throwable
// private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(stack *rtda.OperandStack) {
    _ = stack.PopInt() // dummy
    this := stack.PopRef() // this
    stack.PushRef(this)
    // todo
}

// hack
func jvmgo_SystemOut_println(stack *rtda.OperandStack) {
    str := stack.PopRef()
    this := stack.PopRef()
    this.Class()
    chars := str.Class().GetField("value", "[C").GetValue(str).(*rtc.Obj).Fields().([]uint16)
    for _, char := range chars {
        fmt.Printf("%c", char)
    }
    fmt.Println()
}
