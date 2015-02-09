package misc

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
)

func init() {
    _unsafe(compareAndSwapInt,      "compareAndSwapInt",        "(Ljava/lang/Object;JII)Z")
}

// public final native boolean compareAndSwapInt(Object o, long l, int i, int i1);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(frame *rtda.Frame) {
    stack := frame.OperandStack()
    i1 := stack.PopInt()
    i := stack.PopInt()
    l := stack.PopLong()
    o := stack.PopRef()
    stack.PopRef() // this

    // todo
    fields := o.Fields().([]Any)
    slot := int(l)
    val := fields[slot].(int32)
    if val == i {
        fields[slot] = i1
        stack.PushBoolean(true)
    } else {
        stack.PushBoolean(false)
    }
}
