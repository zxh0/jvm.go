package native

import (
    "fmt"
    "time"
    "unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// register native methods
func init() {
    rtc.SetRegisterNatives(registerNatives)
    system("nanoTime",             "()J",                      nanoTime)
    system("currentTimeMillis",    "()J",                      currentTimeMillis)
    system("identityHashCode",     "(Ljava/lang/Object;)I",    identityHashCode)
    object("getClass",             "()Ljava/lang/Class;",      getClass)

    // hack
    rtc.RegisterNativeMethod("jvmgo/SystemOut", "println", "(Ljava/lang/String;)V", jvmgo_SystemOut_println)
}

func system(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/System", name, desc, method)
}
func object(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/Object", name, desc, method)
}

func registerNatives(operandStack *rtda.OperandStack) {
    // todo
}

// java.lang.System
func nanoTime(stack *rtda.OperandStack) {
    nanoTime := time.Now().UnixNano()
    stack.PushLong(nanoTime)
}
func currentTimeMillis(stack *rtda.OperandStack) {
    millis := time.Now().UnixNano() / 1000
    stack.PushLong(millis)
}
func identityHashCode(stack *rtda.OperandStack) {
    // todo
    ref := stack.PopRef()
    hashCode := int32(uintptr(unsafe.Pointer(ref)))
    stack.PushInt(hashCode)
}

// java.lang.Object
func getClass(stack *rtda.OperandStack) {
    // todo
    panic("obj.getClass()!")
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
