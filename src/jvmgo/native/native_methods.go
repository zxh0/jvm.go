package native

import (
    "fmt"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
    _ "jvmgo/native/java/lang"
    _ "jvmgo/native/java/security"
    _ "jvmgo/native/sun/misc"
    _ "jvmgo/native/sun/reflect"
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
    stack.PopRef() // this
    chars := rtda.JStringChars(str)
    for _, char := range chars {
        fmt.Printf("%c", char)
    }
    fmt.Println()
}
