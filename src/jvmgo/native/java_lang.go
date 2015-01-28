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
    jlSystem("nanoTime",                "()J",                          nanoTime)
    jlSystem("currentTimeMillis",       "()J",                          currentTimeMillis)
    jlSystem("identityHashCode",        "(Ljava/lang/Object;)I",        identityHashCode)
    jlObject("getClass",                "()Ljava/lang/Class;",          getClass)
    jlClass ("getName0",                "()Ljava/lang/String;",         getName0)
    jlClass ("getClassLoader0",         "()Ljava/lang/ClassLoader;",    getClassLoader0)
    jlClass ("desiredAssertionStatus0", "(Ljava/lang/Class;)Z",         desiredAssertionStatus0)
    // hack
    rtc.RegisterNativeMethod("jvmgo/SystemOut", "println", "(Ljava/lang/String;)V", jvmgo_SystemOut_println)
}

func jlSystem(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/System", name, desc, method)
}
func jlObject(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/Object", name, desc, method)
}
func jlClass(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/Class", name, desc, method)
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
    this := stack.PopRef()
    class := this.Class().Obj()
    stack.PushRef(class)
}

// java.lang.Class
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
