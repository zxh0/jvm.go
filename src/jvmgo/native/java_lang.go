package native

import (
    "time"
    "unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// register native methods
func init() {
    system("nanoTime"           ,"()J",                     nanoTime)
    system("currentTimeMillis"  ,"()J",                     currentTimeMillis)
    system("identityHashCode"   ,"(Ljava/lang/Object;)I",   identityHashCode)
    rtc.RegisterNativeMethod("jvmgo/SystemOut", "println", "(Ljava/lang/String;)V", jvmgo_SystemOut_println)
}
func system(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/System", name, desc, method)
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

func jvmgo_SystemOut_println(stack *rtda.OperandStack) {
    panic("totoototototoododod")
}
