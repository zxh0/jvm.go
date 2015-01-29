package native

import (
    "fmt"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
    _ "jvmgo/native/java/lang"
    _ "jvmgo/native/java/security"
)

// register native methods
func init() {
    rtc.SetRegisterNatives(registerNatives)
    // hack!
    rtc.RegisterNativeMethod("jvmgo/SystemOut", "println", "(Ljava/lang/String;)V", jvmgo_SystemOut_println)
}

func registerNatives(frame *rtda.Frame) {
    // todo
}

// hack
func jvmgo_SystemOut_println(frame *rtda.Frame) {
    stack := frame.OperandStack()
    str := stack.PopRef()
    this := stack.PopRef()
    this.Class()
    chars := str.Class().GetField("value", "[C").GetValue(str).(*rtc.Obj).Fields().([]uint16)
    for _, char := range chars {
        fmt.Printf("%c", char)
    }
    fmt.Println()
}
