package lang

import (
    "time"
    "unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _system("nanoTime",             "()J",                      nanoTime)
    _system("currentTimeMillis",    "()J",                      currentTimeMillis)
    _system("identityHashCode",     "(Ljava/lang/Object;)I",    identityHashCode)
}

func _system(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/System", name, desc, method)
}

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
