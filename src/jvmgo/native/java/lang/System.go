package lang

import (
    "time"
    "unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _system("arraycopy",            "(Ljava/lang/Object;ILjava/lang/Object;II)V",   arraycopy)
    _system("currentTimeMillis",    "()J",                                          currentTimeMillis)
    _system("identityHashCode",     "(Ljava/lang/Object;)I",                        identityHashCode)
    _system("nanoTime",             "()J",                                          nanoTime)
}

func _system(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/lang/System", name, desc, method)
}

//arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V
//public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
func arraycopy(frame *rtda.Frame) {
    stack := frame.OperandStack()
    length := stack.PopInt()
    destPos := stack.PopInt()
    dest := stack.PopRef()
    srcPos := stack.PopInt()
    src := stack.PopRef()

    // NullPointerException
    if src == nil || dest == nil {
        panic("NPE") // todo
    }
    // ArrayStoreException
    if !rtc.HaveSameArrayType(src, dest) {
        panic("ArrayStoreException")
    }
    // IndexOutOfBoundsException
    if srcPos < 0 || destPos < 0 || length < 0 ||
            srcPos + length > rtc.ArrayLength(src) ||
            destPos + length > rtc.ArrayLength(dest) {

        panic("IndexOutOfBoundsException") // todo
    }

    rtc.ArrayCopy(src, dest, srcPos, destPos, length)
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
