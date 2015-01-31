package lang

import (
    "time"
    "unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _system("currentTimeMillis",    "()J",                      currentTimeMillis)
    _system("identityHashCode",     "(Ljava/lang/Object;)I",    identityHashCode)
    _system("nanoTime",             "()J",                      nanoTime)
}

func _system(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/System", name, desc, method)
}

func currentTimeMillis(frame *rtda.Frame) {
    stack := frame.OperandStack()
    millis := time.Now().UnixNano() / 1000
    stack.PushLong(millis)
}

// todo
func identityHashCode(frame *rtda.Frame) {
    stack := frame.OperandStack()
    ref := stack.PopRef()
    hashCode := int32(uintptr(unsafe.Pointer(ref)))
    stack.PushInt(hashCode)
}

func nanoTime(frame *rtda.Frame) {
    stack := frame.OperandStack()
    nanoTime := time.Now().UnixNano()
    stack.PushLong(nanoTime)
}
